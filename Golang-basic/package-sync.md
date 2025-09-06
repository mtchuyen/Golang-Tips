## package `sync`

 package sync giúp các hoạt động bất đồng bộ
 
 ***sync.Mutex***: Khóa để xử lý xong tác vụ
 ```go
 m := new(sync.Mutex)
 //some another code
 m.Lock()
 //some action --> another action must stop until thí action finish
 defer m.Unlock()
 ```
 
 ***sync.Atomic***: tương tác với số (integer, float)
 ```go
 atomic.AddInt32(&v, 1)
 ```
***sync.Once***: Khi bạn có một xử lý mà chỉ muốn thực hiện một lần duy nhất

***sync.Cond***: 

***sync.Pool***:
## `sync.Pool`
### Khai báo và sử dụng
```go
var bufPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)  // Chỉ tạo khi pool rỗng
    },
}

func HandleRequest() {
    buf := bufPool.Get().(*bytes.Buffer)  // Lấy buffer cũ từ pool
    defer bufPool.Put(buf)                // Trả lại pool để reuse
    buf.Reset()                          // Xóa data cũ. Quan trọng!!!
    // Sử dụng buffer
}
```
- Tại sao cần `buf.Reset()`? Vì buffer được reuse, nên có thể chứa data từ lần sử dụng trước. `Reset()` xóa sạch data cũ để đảm bảo buffer "clean" cho lần sử dụng mới.

### Cách hoạt động của `sync.Pool`:
```go
// Lần đầu tiên
buf1 := bufPool.Get()  // Pool rỗng → gọi New() → tạo buffer mới

// Sử dụng xong
bufPool.Put(buf1)      // Trả buffer vào pool

// Lần thứ 2
buf2 := bufPool.Get()  // Pool có buffer → lấy buf1 ra dùng lại
// buf2 thực chất là buf1 (cùng memory address)
```
### Lợi ích của sync.Pool là Buffer Memory Reuse
***1. Giảm Memory Allocation***
```
Không reuse: Request1 → Buffer1 → GC
              Request2 → Buffer2 → GC  
              Request3 → Buffer3 → GC

Có reuse:     Request1 → Buffer1 ↓
              Request2 → Buffer1 (reuse) ↓
              Request3 → Buffer1 (reuse) ↓
```

- Cùng một vùng memory được sử dụng cho nhiều request
- Không tạo buffer mới mỗi lần → tiết kiệm memory

***2. Giảm Garbage Collection Pressure***
- Ít object được tạo mới → ít garbage
- GC chạy ít hơn → performance tốt hơn

***3. Better Performance***
- Tăng performance của application. Memory allocation là expensive operation
- Reuse buffer nhanh hơn tạo mới

### Explore Go sync.Pool as Cache
https://medium.com/geekculture/go-sync-pool-as-cache-in-kubernetes-4e247c52e732

## Ref:
- [A Closer Look at Go’s sync Package](https://medium.com/@teivah/a-closer-look-at-go-sync-package-9f4e4a28c35a)
