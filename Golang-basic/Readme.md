## Unit Tests

- [5 Testing Tips in Go](https://medium0.com/star-gazers/5-testing-tips-in-go-3b7f79a546da)

- [How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)

- [Unit Testing CLI Programs in Go](https://medium.com/swlh/unit-testing-cli-programs-in-go-6275c85af2e7)

- [Advanced Testing Tips in Go](https://medium0.com/@mert-akkaya/advanced-testing-tips-in-go-1b4d8eec82a0)

- [Implementing testing in Golang](https://dev.to/lucasnevespereira/implementing-testing-in-golang-4mcp)

- [Unit Testing in Go](https://www.pullrequest.com/blog/unit-testing-in-go/)

- [An Introduction to Testing in Go](https://tutorialedge.net/golang/intro-testing-in-go/)

- [Viết Unit Test trong project với Golang](https://viblo.asia/p/viet-unit-test-trong-project-voi-golang-bWrZnXv95xw)

- [Một số framework dùng cho viết Unit Test cho Golang](https://www.tma.vn/Hoi-dap/Cam-nang-nghe-nghiep/Golang-va-Unit-test/25459)

- [Step by step How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)


## benchmarks

- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [An Introduction to Benchmarking Your Go Programs](https://tutorialedge.net/golang/benchmarking-your-go-programs/)
- [How to write benchmarks in Go](https://progolang.com/how-to-write-benchmarks-in-go/)
- [Analyzing the performance of Go functions with benchmarks](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6)
- [Real Life Go Benchmarking](https://www.cloudbees.com/blog/real-life-go-benchmarking/)
- [Overview of Benchmark Testing in Golang](https://www.geeksforgeeks.org/overview-of-benchmark-testing-in-golang/)

## Map

***Basic and sample:***

- https://www.callicoder.com/golang-maps/
- https://www.geeksforgeeks.org/golang-maps/
- https://www.golangprograms.com/go-language/golang-maps.html
- [Go: Concurrency Access with Maps — Part III](https://medium.com/a-journey-with-go/go-concurrency-access-with-maps-part-iii-8c0a0e4eb27e)

***Advance***:
- https://yourbasic.org/golang/maps-explained/: bài viết có tính học thuật liên quan tới [kỹ thuật hash](https://yourbasic.org/algorithms/hash-tables-explained/)
- [Golang Maps: Not safe for concurrent use](https://golangbyexample.com/go-maps-concurrency/)
- [SOLVED: Golang fatal error: concurrent map writes](https://ashish.one/blogs/fatal-error-concurrent-map-writes/)

#### Check type of map:
***Lib***:
```
import (
	"reflect"
)
```

***Code:***

```
parts = map[int]int{...}
fmt.Printf("%T: %s\n", str, reflect.TypeOf(str).Kind())
```

## Slice

[Go: Slice and Memory Management](https://medium.com/a-journey-with-go/go-slice-and-memory-management-670498bb52be)

[Bad Go: slices of pointers](https://medium.com/@philpearl/bad-go-slices-of-pointers-ed3c06b8bb41)

[Bad Go: not sizing slices](https://medium.com/swlh/bad-go-not-sizing-slices-aed1b01cff83)

## Heap

### [Further Dangers of Large Heaps in Go](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487)

## Structure
***Struct*** (structure) là một `user-defined` data ***type***.

Define một Structure:

```
type StructName struct {
    field1 fieldType1
    field2 fieldType2
}
```

Sau khi `struct` được define, ta có thể khai báo giá trị với syntax:

```
variable_name := structure_variable_type {value1, value2...value_n}
```
Go không support mô hình hướng đối tượng nhưng `structure` gần giống với `class` architecture. Và có thể gọi `Method` (nói ở phần ngay dưới đây) cũng chính là `function`

Ref:
- https://medium.com/@anh.nt094/golang-eb65bfe1a8bb


## Method

***Method*** cũng chính là ***function***, nhưng nó thuộc về 1 ***type*** nhất định. Cách define của method hơi ***khác*** với function. Khác ở chỗ nó cần thêm một parameter gọi là ***receiver*** — chính là ***type*** được nhắc tới ở trên. Với cách define này, method có thể access vào các properties của receiver (các field của struct).

Syntax:
```
func (r Type) methodName() [return_type]{
   /* method body*/
}
```
- ***(r Type)*** là parameter khác biệt so với khai báo 1 method với 1 function, parameter này được gọi là ***receiver***.
- Trong một package, có thể có nhiều method trùng tên, function KHÔNG thể trùng tên. Ta có thể tạo nhiều method trùng tên nhưng phải khác ***receiver***. Ví dụ như cùng là method `DienTich` (để tính diện tích), có thể tên giống nhau ở 2 ***receiver***: Circle & HinhVuong.

Ví dụ:

```go
type Circle struct {
	bankinh float64
}

func (c Circle) DienTich() float64 {
	return math.Pi * c.bankinh * c.bankinh
}

```
- ***Circle*** là 1 type.
- `bankinh` là 1 property của type ***Circle***.
- `DienTich()` là một `method` của `type` ***Circle***, và `receiver` khai báo là `c Circle`  trong khai báo của method.


### Method declaration
- Define: https://golang.org/ref/spec#Method_declarations
- [Method declaration with function receiver in Golang](https://pgillich.medium.com/method-declaration-with-function-receiver-in-golang-7f5531ded97d)


## [Interface](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/interface.md)

## Pipeline
- [Go Pipeline for a layman](https://anupamgogoi.medium.com/go-pipeline-for-a-layman-4791fb4f1e2d)
- 

## Mutex
- [Mutex Examples in Go](https://levelup.gitconnected.com/mutex-examples-in-go-ad7c440461a4)
- 
