
***Note***: tại sao phải đưa bài viết này vào github, bởi vì medium hạn chế lượng đọc bài viết, khiến cho việc re-read gặp khó khăn.

## Các tiêu chí so sánh thư viện cache:

- High concurrency
- Data structure
- Capacity & evict policy
- ttl
- Stats statistics
- Standalone

Có được tiêu chí so sánh này thì sẽ biết lựa chọn các cache-lib phù hợp với dự án
- UC1: Dự án cần cache không bị out-of-date thì cần có `ttl`
- UC2: Dự án cần số lượng cache lớn thì cần capacity.
- UC3: Dự án có cache-item phức tạp thì cần data-structure tổ chức tốt
- ...


### 1. High concurrency

Generally, the cache uses ***locking*** in concurrency control. Assuming that there are 100 elements and only one read-write lock, then you have to do it one by one when two elements are to be written at the same time. Low efficiency, right?

So ***buckets*** are introduced, dividing the 100 elements into 10 groups, and controlling the 10 buckets with ten read-write locks respectively. In this case, there is only a 1% chance those two elements are in the same bucket at the same time, thus 100 elements can probably write simultaneously, improving the concurrency efficiency by 100 times!

![buckets-shard](https://github.com/mtchuyen/Golang-Tips/blob/master/statics/cache_2_6UxZSfU4FVJDnTwNxaNLtw.png)

Như trong hình, cache được chia làm 10 (`shard_num`) buckets (10 group, 10 shard...) theo 1 giải thuật nào đó tùy chọn.

### 2.Data structure
The memories differ between caches, and we will mainly focus on:
- Memory overhead. Whether the data structure design is reasonable, and whether additional memory allocation is required to implement some functions.
- Key and value store efficiency.
- GC influence.

Almost all caches depend on `map` for storage. Therefore, improving the efficiency of `map` and avoiding additional `GC` overhead can make the cache much more efficient.

There are usually several methods.
- Allocate memory outside of the stack.
- Don’t use pointers in `map`, but only use primitive types like `map[int]string`. For pointers are often allocated to the heap, not using them is an effective way to stop frequent GC.
- Use simple data types and avoid complex ones to make the map as small as possible. If you are interested, please read [runtime: Large maps cause significant GC pauses](https://github.com/golang/go/issues/9477).

### 3. Capacity & evict policy

The scalability of the cache’s capacity determines the evict policies. The common evict policy algorithms include FIFO, LRU, etc.

Capacity của 1 cache là lượng item có thể lưu trữ trong cache được.

Capacity phụ thuộc vào yếu tố `kích thước data` của item. Vì vậy với sự giới hạn của memory app, thì Capacity của các thư viện cache cần được đánh giá thêm với các điều kiện:
- Cache item có trong memory
- Khả năng flush data từ memory xuống disk

Để đảm bảo app có thể chạy tốt (không bị tràn memory) khi cache overload, thì `Capacity` và `evict policy` cần được cân bằng với nhau
- Nếu `Capacity` là `unlimited` thì cần có `evict policy` (time-to-live)
- Nếu `Capacity` là `limited` thì KHÔNG cần có `evict policy`, vì bản thân `length-of-cache` luôn cố định rồi.

### 4. Time-to-live

***ttl***, a feature pursued by many cache users, can save the users from worrying about cache explosion in many cases. And you need to adjust its value according to statistics only when performance issues occur.


### 5. Statistics
Whether the stats, including the number of cache hits and the number of requests, are exposed is also what we should consider when evaluating caches.

Các chỉ số thống kê này khá quan trọng với bước đầu sử dụng cache-lib, ví dụ:
- nếu có chỉ số này sẽ biết được tỉ lệ hit/miss để quyết định kích thước cache, loại cache,...
- biết được delmiss, delhits để tăng kích thước `ttl`, customize các chỉ số khởi tạo cache phù hợp

### 6. Standalone

If a cache can support separate deployment, then its tuning, deployment, and performance monitoring become easier.

Có thể nói về tiêu chí này:
- Triển khai độc lập (không tích hợp vào app)
- Mở rộng nhiều cổng giao tiếp (interface)

Thực sự thì tiêu chí này đưa vào cũng không thể đánh giá được sự tiện lợi của cache-lib, cái này được nói là tính năng của cache-lib thì đúng hơn. Việc đánh giá sử dụng tiêu chí này còn phụ thuộc vào mục tiêu của ứng dụng.
> Nhiều tính năng quá nhiều khi (dùng không hết) khiến app nặng nề, hoạt động chậm chạp.



## So sánh 4 thư viện cache: `go-cache`, `bigcache`, `golang-lru`, & `groupcache`.

https://medium.com/codex/our-go-cache-library-choices-406f2662d6b

Note: Bảng so sánh này chủ yếu dùng với các ***version < 1.18***. 

***Bảng so sánh dựa trên các tiêu chí***

![cachelib-compare](https://github.com/mtchuyen/Golang-Tips/blob/master/statics/cache_1_5cBSAXhaw9LCv76DIdIeuQ.png)

***Bảng so sánh benchmark***

![cache-benchmark](https://github.com/mtchuyen/Golang-Tips/blob/master/statics/cache_4_4WBVLam17EBz1Kya.png)

- ***bigcache*** shows no advantages in either `Get` or `Put`. Its GC effect is even worse than ***go-cache*** and ***groupcache***.
- ***groupcache*** and ***golang-lru*** don’t support sharding, but they are the most efficient.

Kết quả này có thể gây ra sự xung đột (hiểu nhầm), tranh cãi về việc cache-lib nào hiệu quả. Nhưng nhìn vào cách benchmark có thể thấy đây chỉ là 1 kết quả nhất định với 1 dạng `key-value` nhất định, thực tế (production environment) triển khai app mới đánh giá đúng hiệu năng của từng loại cache-lib.


### 1. High concurrency

Only ***go-cache*** and ***bigcache*** support sharding, both supporting a configurable number of buckets and controlling the concurrency of each bucket by `sync.RWMutex`. 

Their difference is that go-cache applies [*djb2*](https://github.com/patrickmn/go-cache/blob/46f407853014144407b6c2ec7ccc76bf67958d93/sharded.go#L34) hash algorithm, while bigcache uses [*FNV*](https://github.com/allegro/bigcache/blob/77bbae442c22ec7ba79f77cabd56d0e5a9366109/fnv.go#L20). I haven’t done much research on the hash algorithm but can see these two seem not much difference in speed from this StackExchange [post](https://softwareengineering.stackexchange.com/questions/49550/which-hashing-algorithm-is-best-for-uniqueness-and-speed).

In contrast, neither ***groupcache*** nor ***golang-lru*** support sharding.

### 2.Data structure
***groupcache*** and ***golang-lru*** are exactly the same that they both save the key and value using the `List.Element` in `container/list` , and then save `<key, list.Element>` with `map`, and save the element to a `List` to provide the function of LRU.

For ***go-cache***, there is nothing special in its data structure, which uses the simple `map[string]Item` , in which `Item` is a struct containing an `interface{}` and expire. When the object is too large, go-cache’s performance might be affected by GC.

***bigcache*** is the most distinctive and GC-friendly cache. It designs an underlying storage structure that separates the value from the [hashmap](https://github.com/allegro/bigcache/blob/77bbae442c22ec7ba79f77cabd56d0e5a9366109/shard.go#L19), and stores the value in a [BytesQueue](https://github.com/allegro/bigcache/blob/77bbae442c22ec7ba79f77cabd56d0e5a9366109/shard.go#L20), while the `map` only stores the simplest data structure `map[uint64]uint32`, within which the value is the index position in the queue.

![bigcache-datastrucure](https://github.com/mtchuyen/Golang-Tips/blob/master/statics/cache_3_vMGCvMLINna4o1gakFXpgw.png)

Có thể thấy rằng ***BigCache*** tổ chức data structure khéo léo (phức tạp) nhất trong 4 loại cache trên.

### 3. Capacity & evict policy

As we all know, neither ***go-cache*** nor ***bigcache*** sets limits on the number of keys, indefinite addition is allowed, so ***no*** evict policy is needed. But the other two have limitations.

***golang-lru***, only supports a fixed number of keys and the LRU algorithm to eliminate old data. It saves the old data by a double-linked list, and each time it deletes the oldest.

***groupcache*** provides the code for limiting the number of keys, but without a default number, and sets a limit to the total data size, which means the number is unlimited if size allows. Meanwhile, groupcache also eliminates the old data with the double-linked list of the `container/list` package and the LRU algorithm combined.

### 4. Time-to-live

***golang-lru*** and ***groupcache*** are *not* `ttl` supportive, so you have to clear the data manually, or expect the automatic delete of the earliest data when the cache is full. But cache miss will be caused once the cache is inserted frequently.

Trong bảng trên ta thấy ***golang-lru*** & ***groupcache*** đều có có `capacity-limited` nên tính năng `ttl` với lived-of-item không cần thiết. Item sẽ bị loại bỏ (evict) khi `size of cache` tới hạn (limited). Cơ chế loại bỏ (evict policy) đã nói ở trên. Nếu sử dụng thư viện nào thì cần check xem thư viện đó sử dụng hỗ trợ policy nào, hoặc app cấu hình với policy đó theo nhu cầu.

***bigcache*** supports all keys to be with a unified `ttl`, which is 10 minutes by default. There are two ways to remove the expired keys.
- Start a 5-minute timer to clean up periodically.
- Determine whether the `oldest` key needs to be deleted every time when `Set` is called.

Do đó cũng có 2 phương thức để xóa expire item:
- `Janitor`, automatically triggers the cleaning method with `timer`, monitors each key’s expiration time circularly, and deletes the key when it expires.
- Check if the key is expired when calling `Get`, if so, return nil and delete.

### 5. Statistics

Both ***groupcache*** and ***bigcache*** support stats statistics, while the other two don’t.
As to implementation, ***groupcache*** only needs a CacheStats to count hits or gets each time `Get` is called, since it does not support bucketing. But ***bigcache*** needs a `hashmapStats` map for statistics in each cache, and exposes hits and misses data through a Stats. 


### 6. Standalone

Only ***groupcache*** and ***bigcache*** support separate deployment and provide related HTTP interfaces for external access.

***bigcache*** can expose methods like `GET`, `PUT`, and `DELETE` by starting an httpServer.

***groupcache*** claims to be the replacement of `memcached`, supporting the requests through the ***proto buffer protocol*** as well as httpServer.


## generics-cache
`Go generic cache` được cho là sẽ cải thiện performance của Cache, và support đầy đủ các tính năng (so với 4 thư viện cache đã so sánh ở trên)

***generic-cache*** uses ***comparable*** interface as key type, and supports all kinds of cache algorithms, e.g. FIFO, LRU, LFU, MRU, etc.

***Một số đặc điểm:***
- Implementing simple concurrency control, but no sharding solution provided.
- One key, one `ttl`, and Set a timer by starting a goroutine to trigger the cleaning. It is maybe the potential bottleneck if the cache capacity reaches a magnitude.
- Building the cache through the `container/ring` data structure.
- All cache implementations handle values with `pointers`, which are `not` GC friendly.

***Một số thư viện hỗ trợ generic-cache:***

- https://github.com/Code-Hex/go-generics-cache
- https://github.com/dboslee/lru
- [Memory Cache: A zero-dependency cache library for storing data in memory with generics.](https://github.com/rodrigo-brito/memory-cache)
- [Cache is a thread safe, generic, and sharded in memory LRU cache object](https://github.com/alistanis/cache)
- [Thread safe, generic, in-memory cache for Golang with optional TTL settings](https://github.com/twharmon/gocache)
- [ECSGo is an Entity Component System(ECS) in Go: Cache friendly, Multi threading](https://github.com/kongbong/ecsgo)


Note: Những thư viện này dùng `generic` của Go, nên phải dùng các version ***> 1.18***.

## Ref
- https://dgraph.io/blog/post/caching-in-go/
- https://medium.com/codex/our-go-cache-library-choices-406f2662d6b

