
Note: tại sao phải đưa bài viết này vào github, bởi vì medium hạn chế lượng đọc bài viết, khiến cho việc re-read gặp khó khăn.

## Các tiêu chí so sánh thư viện cache:

- High concurrency
- Data structure
- Capacity & evict policy
- ttl
- Stats statistics
- Standalone

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


## So sánh 4 thư viện cache: `go-cache`, `bigcache`, `golang-lru`, & `groupcache`.

https://medium.com/codex/our-go-cache-library-choices-406f2662d6b

Note: Bảng so sánh này chủ yếu dùng với các ***version < 1.18***. 

![cachelib-compare](https://github.com/mtchuyen/Golang-Tips/blob/master/statics/cache_1_5cBSAXhaw9LCv76DIdIeuQ.png)

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
