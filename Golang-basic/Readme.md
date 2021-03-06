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

## Interface
https://research.swtch.com/interfaces


## Slice

[Go: Slice and Memory Management](https://medium.com/a-journey-with-go/go-slice-and-memory-management-670498bb52be)

[Bad Go: slices of pointers](https://medium.com/@philpearl/bad-go-slices-of-pointers-ed3c06b8bb41)

[Bad Go: not sizing slices](https://medium.com/swlh/bad-go-not-sizing-slices-aed1b01cff83)

## Heap

### [Further Dangers of Large Heaps in Go](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487)

## Method

### Method declaration
- Define: https://golang.org/ref/spec#Method_declarations
- [Method declaration with function receiver in Golang](https://pgillich.medium.com/method-declaration-with-function-receiver-in-golang-7f5531ded97d)
- 
