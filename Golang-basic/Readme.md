## Unit Tests
- https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Unit-test.md

## benchmarks

- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [An Introduction to Benchmarking Your Go Programs](https://tutorialedge.net/golang/benchmarking-your-go-programs/)
- [How to write benchmarks in Go](https://progolang.com/how-to-write-benchmarks-in-go/)
- [Analyzing the performance of Go functions with benchmarks](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6)
- [Real Life Go Benchmarking](https://www.cloudbees.com/blog/real-life-go-benchmarking/)
- [Overview of Benchmark Testing in Golang](https://www.geeksforgeeks.org/overview-of-benchmark-testing-in-golang/)

## [Các data Type trong Go](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/DataType.md)

Golang có các data ***Type***:
- ***Basic*** (Cơ bản) Types: Numbers (int32, int64, float,...), Booleans (true or false), Strings
- ***Aggregate*** (Tổng hợp) Types: Structs, Arrays
- ***Reference*** (Tham chiếu) Types: Pointers, Slices, Maps, Functions, Channels
- ***Interface*** Type:


- https://kyxey.medium.com/lets-go-part-6-basic-types-5d66208a63c9
- https://kyxey.medium.com/lets-go-part-7-aggregate-types-1-structs-406a7d81c813
- https://kyxey.medium.com/lets-go-part-9-reference-types-1-pointers-8ffc5ad69b7e
- https://kyxey.medium.com/lets-go-part-10-reference-types-2-slices-18224077c1aa
- https://kyxey.medium.com/lets-go-part-11-reference-types-3-maps-849f4e3c91bb
- https://kyxey.medium.com/lets-go-part-12-reference-types-4-channels-555c2aad4c28


### Map

***Basic and sample:***

- https://www.callicoder.com/golang-maps/
- https://www.geeksforgeeks.org/golang-maps/
- https://www.golangprograms.com/go-language/golang-maps.html
- [Go: Concurrency Access with Maps — Part III](https://medium.com/a-journey-with-go/go-concurrency-access-with-maps-part-iii-8c0a0e4eb27e)
- [Avoid The Three Mistakes When Using a Map in Go](https://levelup.gitconnected.com/avoid-the-three-mistakes-when-using-a-map-in-go-699435db226c)
- [Go maps: declaring and initializing](https://bitfieldconsulting.com/golang/map-declaring-initializing)

***Advance***:
- https://yourbasic.org/golang/maps-explained/: bài viết có tính học thuật liên quan tới [kỹ thuật hash](https://yourbasic.org/algorithms/hash-tables-explained/)
- [Golang Maps: Not safe for concurrent use](https://golangbyexample.com/go-maps-concurrency/)
- [SOLVED: Golang fatal error: concurrent map writes](https://ashish.one/blogs/fatal-error-concurrent-map-writes/)
- [map[string]interface{} in Go](https://bitfieldconsulting.com/golang/map-string-interface)
- [Finding whether a Go map key exists](https://bitfieldconsulting.com/golang/map-key-exists)


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

### [Slice](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Readme.md#slice)

### [Structure](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Readme.md#structure)

## Method

***Method*** cũng chính là ***function***, nhưng nó thuộc về 1 ***type*** nhất định. Cách define của method hơi ***khác*** với function. Khác ở chỗ nó cần thêm một parameter gọi là ***receiver*** — chính là ***type*** được nhắc tới ở trên. Với cách define này, method có thể access vào các properties của receiver (các field của struct).

[See more here](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Method.md)

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


## [Data Structure](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/DataStructure.md)

Phân biệt  khái niệm “data structure” trong nghĩa ***logic*** và “data type category” trong ***ngữ pháp Go*** (the language spec).

| Khía cạnh       | **Data Structure**                                   | **Data Type Category (theo Go spec)**                                    |
| --------------- | ---------------------------------------------------- | ------------------------------------------------------------------------ |
| Mục tiêu        | Mô hình tổ chức và lưu trữ dữ liệu (logic, abstract) | Cách Go **phân loại kiểu dữ liệu** theo **cơ chế lưu trữ và tham chiếu** |
| Ví dụ           | Stack, Queue, Set, List, Tree, Map...                | Basic, Aggregate, Reference, Interface                                   |
| Tính trừu tượng | Trừu tượng hoá logic thuật toán                      | Trừu tượng hoá cách hoạt động bộ nhớ                                     |
| Xuất hiện ở đâu | Thuộc **Computer Science**                           | Được định nghĩa trong **Go Language Specification**                      |

Note:
- Tất cả *Reference types, Aggregate types* (của `Data Type`) đều là nền tảng để xây dựng `Data Structures`.
- Go chia loại theo cách ***hoạt động vùng nhớ***, không phải theo công dụng logic.
- **Data structure** = khái niệm logic → mọi `Aggregate` hoặc `Reference` type đều có thể dùng để xây dựng `data structure`.
- `Function`, `Pointer`, `Channel` là `Reference` type, nhưng **không phải** `data structure` vì không biểu diễn cấu trúc dữ liệu logic (chúng là “cơ chế”, không phải “mô hình dữ liệu”).

| Loại                  | Nhóm theo Go spec | Kiểu truyền                 | Có thể xem là Data Structure? | Giải thích                                    |
| --------------------- | ----------------- | --------------------------- | ----------------------------- | --------------------------------------------- |
| **int, string, bool** | Basic             | Value                       | ❌                             | Primitive, không chứa cấu trúc                |
| **array**             | Aggregate         | Value                       | ✅                             | Tập hợp cố định phần tử                       |
| **struct**            | Aggregate         | Value                       | ✅                             | Gom nhiều trường dữ liệu                      |
| **slice**             | Reference         | Reference                   | ✅                             | Trỏ đến vùng mảng thực                        |
| **map**               | Reference         | Reference                   | ✅                             | Hash table nội bộ                             |
| **channel**           | Reference         | Reference                   | ⚙️ Gián tiếp                  | Dùng để giao tiếp, không tổ chức dữ liệu      |
| **pointer**           | Reference         | Reference                   | ❌                             | Trỏ 1 địa chỉ đơn                             |
| **function**          | Reference         | Reference                   | ❌                             | Lưu hành vi, không lưu dữ liệu                |
| **interface**         | Interface         | Special (type + value pair) | ⚙️ Không hẳn                  | Trừu tượng hoá hành vi, không tổ chức dữ liệu |

## [Interface](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/interface.md)

## Pointers

[The Fundamentals of Pointers in Go](https://betterprogramming.pub/pointers-in-go-9aa5c0682a)

[When and Where to Use Pointers in Go](https://towardsdev.com/when-and-where-to-use-pointers-in-go-7e89643a2c9)

[Golang: nắm chắc khái niệm CON TRỎ trong 5 phút](https://viblo.asia/p/go-lang-nam-chac-khai-niem-con-tro-trong-5-phut-maGK7OXAKj2)

[Series Golang cơ bản (Phần 15: Con trỏ - Pointer)](https://techmaster.vn/posts/35009/series-golang-co-ban-phan-15-con-tro-pointer)

[Con trỏ trong Go](https://viblo.asia/p/con-tro-trong-go-3P0lP6mPKox)

[Go – Con trỏ](https://phocode.com/go/go-lap-trinh-go/go-con-tro/)

[Con trỏ](https://huynhphuchuy.com/golang-series/bai-10-kieu-du-lieu-trong-golang-con-tro/)



## Pipeline
- [Go Pipeline for a layman](https://anupamgogoi.medium.com/go-pipeline-for-a-layman-4791fb4f1e2d)
- 

## [Mutex](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Mutex.md)


## Heap

### [Further Dangers of Large Heaps in Go](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487)

