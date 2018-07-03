
# Goroutine
Bài viết:
[Concurrency and Parallelism in Golang](https://medium.com/@tilaklodha/concurrency-and-parallelism-in-golang-5333e9a4ba64)

Tìm hiểu về Concurrent và Parallel:

- Định nghĩa và sự khác biệt

- Kiến trúc từng thành phần

[Link1](https://kipalog.com/posts/7-concurrency-models-in-seven-week--phan-1)

[Link2](http://thachleblog.com/phan-biet-parallelism-va-concurrency/)

[Do not communicate by sharing memory. Instead, share memory by communicating](http://www.minaandrawos.com/2015/12/06/concurrency-in-golang/)


[Concurrency Programming Guide](https://viblo.asia/p/concurrency-programming-guide-63vKjpYdl2R)


### Channel

Channel là tính năng cho phép 2 goroutine giao tiếp/trao đổi dữ liệu với nhau.

***Điều hướng channel***

[Source](http://phocode.com/go/go-lap-trinh-go/go-concurrency/)

Chúng ta có thể quy định channel chỉ được phép đọc hoặc chỉ được phép ghi dữ liệu vào đó. Ví dụ:

> func pinger(c ***chan<-*** string)

Dòng code trên quy định channel c chỉ được truyền ***dữ liệu vào***.

> func printer(c ***<-chan*** string)

Dòng code trên quy định channel c chỉ được phép ***đọc dữ liệu***.

Nếu không quy định hướng đi của channel thì mặc định channel sẽ được phép vừa đọc vừa ghi.


***1. Check number goroutine is running***

```go
type CountWG struct {
  sync.WaitGroup
  Count int // Race conditions, only for info logging.
}
// Increment on Add
func (cg CountWG) Add(delta int) {
  atomic.AddInt32(&cg.Count, delta)
  cg.WaitGroup.Add(delta)
}
// Decrement on Done
func (cg CountWG) Done() {
  atomic.AddInt32(&cg.Count, -1)
  cg.WaitGroup.Done()
}
```
