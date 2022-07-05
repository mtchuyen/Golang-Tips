## Concurrency
http://runikitkat.com/post/go-routine-under-the-hood/

***Concurrency***:
- Concurrency trong một chương trình là khi chúng ta cho phép chạy nhiều hơn một công việc (task) một cách đồng thời.
- Concurrency không phải là Parallelism
- Cocurrency đi kèm hai khái niệm cơ bản: goroutines và channels

***Goroutine***:
- Một goroutine là một hàm mà có thể chạy *đồng thời* với các hàm khác kể cả hàm *main*, nhưng lại phụ thuộc hàm *main*: nếu *main* kết thúc thì hàm đó cũng bị cưỡng bức kết thúc (terminate).
- Goroutines được xem như như thread nhưng nhẹ hơn, tuy nhiên nó không phải là một tiến trình(process) hay là thread của hệ thống(OS).
- Lý thuyết hoạt động goroutines dựa trên sự chia sẻ vùng nhớ.
- Goroutines rất rẻ. Một goroutine được tạo ra chỉ tốn 2KB trong stack, và khi chạy xong bị huỷ bởi runtime. Chúng ta có thể sử dụng goroutines thoải mái mà không phải lo nghĩ về việc tốn kém bộ nhớ. 

***Channels***:
- Channel sinh ra dùng để giao tiếp giữa 2 goroutines, bao gồm gửi và nhận dữ liệu.
- Channel là reference type.
- Về cơ bản, concept của channel là “typed pipes”. Nó tạo một đường ống liên kết giữa 2 goroutines, chúng ta có thể gửi các object phức tạp qua channel.
- Channel có thể dùng cho synchronization.

### Why doesn’t Go Support Map Concurrency at the Language Level?

https://medium.com/@ClaytonXia/why-doesnt-go-support-map-concurrency-at-the-language-level-17b715587163

#### Why is it not natively supported? The official answer is listed below:
- ***Typical usage scenario:** A typical usage scenario for map is that safe access from multiple goroutines is not required.
- ***Atypical scenarios:*** The map may be part of some larger data structure or an already synchronized computation.
- ***Performance scenario considerations:*** If only a few programs are added to the security, all operations of the map have to deal with mutex, which will reduce the performance of most programs.

So after a long discussion, the Go team believes that native maps should be more suitable for typical usage scenarios.

If it is for a small number of cases, it will cause most programs to pay an extra performance cost , and it is decided not to support native concurrent map reading and writing. And since Go1.6, a detection mechanism has been added , and concurrency will lead to exceptions.

#### Why crash
As mentioned earlier, the concurrency detection of native maps will be performed since Go1.6, which is a “nightmare” for some people. Some people may complain: “It’s okay to return a mistake, why simply make my program crash”.

Let’s assume that the concurrent read and write of map goes into the following two scenarios:

***Generate panic:*** program panic -> leads to recover -> does not process concurrent map -> map has dirty data -> program uses dirty data -> generates unknown risks

***Generate crash:*** program crash -> crash directly -> save data (data is normal) -> generate known risks

Which option would you choose? Go officials chose the second one after assessing the cost.

#### Let it crash

The method chosen by the Go official team is the classic “let it crash” behavior in the industry, which is adopted as a design philosophy in many programming languages.

‘Let it crash’ lets engineers don’t have to worry too much about unknown bugs and instead do comprehensive defensive coding. Erlang is well known for this concept.

The reason of inability to read and write the map concurrently directly is the consideration of “let it crash”. If you want to avoid this situation in your own project, you can add race detection (-race) to tool chains such as linter, and you can also avoid such risks.

How do you like the design considerations for Go? Comment or discussions welcomed~

## Goroutine:
### Khai báo:

Sử dụng bằng cách thêm keyword `go` trước một hàm. 
```go
go say("hello")
```
Chúng ta có thể define số goroutines chạy cùng lúc tối đa bằng khai báo:
```
export GOMAXPROCS=100
```
### Sử dụng:

## Channel
### Khai báo:

Chúng ta tạo channel bằng ***make***:
```go
chInt := make(chan int)
chPully := make(chan Pully) // Pully là interface
```
### Sử dụng:
-  Để gửi data thông qua channel:
```go
chInt <- 3
```
- Để nhận data:
```go
x := chanInt
or 
x := <- ch
```
- Kiểm tra channel đóng
```go
_, ok = <- c // ok bằng true nếu c còn mở
```
- Close một channel
```go
close(c)
```

## Bài viết:
[Concurrency and Parallelism in Golang](https://medium.com/@tilaklodha/concurrency-and-parallelism-in-golang-5333e9a4ba64)

Tìm hiểu về Concurrent và Parallel:

- Định nghĩa và sự khác biệt

- Kiến trúc từng thành phần

[Link1](https://kipalog.com/posts/7-concurrency-models-in-seven-week--phan-1)

[Link2](http://thachleblog.com/phan-biet-parallelism-va-concurrency/)

[Do not communicate by sharing memory. Instead, share memory by communicating](http://www.minaandrawos.com/2015/12/06/concurrency-in-golang/)


[Concurrency Programming Guide](https://viblo.asia/p/concurrency-programming-guide-63vKjpYdl2R)

[Pool Go Routines To Process Task Oriented Work](https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html)

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
## Managing Groups of Goroutines in Go
[Managing Groups of Goroutines in Go](https://medium.com/swlh/managing-groups-of-gorutines-in-go-ee7523e3eaca)

### Detecting Goroutine Leaks with Test Cases
https://engineering.razorpay.com/detecting-goroutine-leaks-with-test-cases-b0f8f8a88648

### Common Goroutine Leaks that You Should Avoid
https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee


# Referral:
- Lê Ngọc Thạch, [Concurrency trong Go Lang](https://devblog.dwarvesf.com/post/concurrency/)
