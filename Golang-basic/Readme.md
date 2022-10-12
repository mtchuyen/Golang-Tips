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

- [Golang — basic unit testing and benchmarking](https://sherryalex29.medium.com/golang-basic-unit-testing-and-benchmarking-10ce064f1081)

## benchmarks

- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [An Introduction to Benchmarking Your Go Programs](https://tutorialedge.net/golang/benchmarking-your-go-programs/)
- [How to write benchmarks in Go](https://progolang.com/how-to-write-benchmarks-in-go/)
- [Analyzing the performance of Go functions with benchmarks](https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6)
- [Real Life Go Benchmarking](https://www.cloudbees.com/blog/real-life-go-benchmarking/)
- [Overview of Benchmark Testing in Golang](https://www.geeksforgeeks.org/overview-of-benchmark-testing-in-golang/)

## Các data Type trong Go
Golang có các data ***Type***:
- ***Basic*** (Cơ bản) Types: Numbers (int32, int64, float,...), Booleans (true or false), Strings
- ***Aggregate*** (Tổng hợp) Types: Structs, Arrays
- ***Reference*** (Tham chiếu) Types: Pointers, Slices, Maps, Functions, Channels
- ***Interface*** Type:

### Kiểu basic:

Tham khảo: https://github.com/buihien0109/golang-basic

### 1. Kiểu bool

Bool là kiểu dữ liệu chỉ nhận 2 giá trị hoặc true hoặc false (hoặc đúng hoặc sai).

```golang
a := true //a được gán bằng true
b := false // b được gán bằng false
fmt.Println("a:", a, "b:", b) 

c := a && b // c được gán bằng a&&b 
fmt.Println("c:", c) 

d := a || b // d được gán bằng a||b
fmt.Println("d:", d) 
```
---
### 2. Kiểu dữ liệu số (numeric type)
#### Số nguyên
Các kiểu số nguyên trong Go là uint8, uint16, uint32, uint64, int8, int16, int32, int64, int. 
> **uint** tức là **unsigned int** – là kiểu số nguyên không âm

| Kiểu   | Giới hạn                                   |
|--------|--------------------------------------------|
| uint8  | 0 – 255                                    |
| uint16 | 0 – 65535                                  |
| uint32 | 0 – 4294967295                             |
| uint64 | 0 – 18446744073709551615                   |
| int8   | -128 – 127                                 |
| int16  | -32768 – 32767                             |
| int32  | -2147483648 – 2147483647                   |
| int64  | -9223372036854775808 – 9223372036854775807 |

```golang
var num3 int32 = 20132
var num4 int32 = 23244
fmt.Println("Tong 2 so num3 và num4 là:", num3 + num4)

var num5 int = 20132
var num6 int = 23244	
fmt.Println("Tong 2 so num5 và num6 là:", num5 + num6)
```

#### Số thực
Trong Go có 2 kiểu số thực là **float32** và **float64**.

Thông thường để biểu diễn số thực, bạn chỉ cần dùng **float64** là đủ.

```golang
a, b := 5.67, 8.97
fmt.Printf("Kiểu dữ liệu của a là %T và của b là %T\n", a, b)

sum := a + b
sub := a - b

fmt.Println("Tổng a và b:", sum)
fmt.Println("Hiệu a và b:", sub)
```

#### Số phức
Trong Go, có 2 kiểu số phức là **complex64** và **complex128**

- Complex64: Số phức có phần thực **float32** và phần ảo
- Complex128: Số phức có phần thực **float64** và phần ảo

```golang
c1 := complex(5, 7)
c2 := 8 + 27i

cadd := c1 + c2
fmt.Println("Tổng 2 số phức c1 và c2:", cadd)

cmul := c1 * c2
fmt.Println("Tích 2 số phức c1 và c2:", cmul)
```


#### Các kiểu dữ liệu số khác

- **byte** là tên gọi khác của **uint8**
- **rune** là tên gọi khác của **int32**

---
### Kiểu dữ kiệu chuỗi
Trong Go, chuỗi là một tập hợp các byte

```golang
first := "Bùi"
last := "Hiên"

name := first + " " + last
fmt.Println("My name is ",name)
```

### Ép kiểu
Ngôn ngữ Go rất nghiêm ngặt và chặt chẽ, nên chúng không cho phép tự động chuyển đổi (ép kiểu) kiểu dữ liệu.
```golang
//WRONG
i := 55      //int
j := 67.8    //float64
sum := i + j //int + float64 not allowed
fmt.Println(sum)
```

> Sửa lại thành
```golang
//CORRECT
i := 55      //int
j := 67.8    //float64
sum := i + int(j) //j is được ép kiểu thành int
fmt.Println(sum)
```

- https://kyxey.medium.com/lets-go-part-6-basic-types-5d66208a63c9
- https://kyxey.medium.com/lets-go-part-7-aggregate-types-1-structs-406a7d81c813
- https://kyxey.medium.com/lets-go-part-9-reference-types-1-pointers-8ffc5ad69b7e
- https://kyxey.medium.com/lets-go-part-10-reference-types-2-slices-18224077c1aa
- https://kyxey.medium.com/lets-go-part-11-reference-types-3-maps-849f4e3c91bb
- https://kyxey.medium.com/lets-go-part-12-reference-types-4-channels-555c2aad4c28


## Map

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

## Slice

[Go: Slice and Memory Management](https://medium.com/a-journey-with-go/go-slice-and-memory-management-670498bb52be)

[Bad Go: slices of pointers](https://medium.com/@philpearl/bad-go-slices-of-pointers-ed3c06b8bb41)

[Bad Go: not sizing slices](https://medium.com/swlh/bad-go-not-sizing-slices-aed1b01cff83)


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

### Struct Composition in Go
https://medium.com/@muhammadarifineffendi/struct-composition-in-go-80492bd447bd

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

## Mutex
- [Mutex Examples in Go](https://levelup.gitconnected.com/mutex-examples-in-go-ad7c440461a4)
- 

## Heap

### [Further Dangers of Large Heaps in Go](https://syslog.ravelin.com/further-dangers-of-large-heaps-in-go-7a267b57d487)

