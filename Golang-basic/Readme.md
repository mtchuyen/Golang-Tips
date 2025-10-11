## Unit Tests
- https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Unit-test.md

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
### 3. Kiểu dữ kiệu chuỗi

#### Tính bất biến (immutability) của Strings
Strings trong Go gặp vấn đề lớn nhất là ***tính bất biến (immutability)***. Điều này có nghĩa là một khi một chuỗi đã được tạo, bạn không thể thay đổi nội dung của nó.

Khi bạn thực hiện các thao tác tưởng chừng như là "thay đổi" một chuỗi, ví dụ như nối chuỗi (concatenation), Go sẽ không sửa chuỗi ban đầu mà thay vào đó sẽ tạo ra một chuỗi hoàn toàn mới trong bộ nhớ.

***Vấn đề của String bất biến***

Hãy xem xét ví dụ sau:

```Go
s := "hello"
s += " world"
```
Khi bạn chạy đoạn code trên, Go thực hiện các bước sau:
```
1. Tạo chuỗi s với giá trị "hello" trong bộ nhớ.

2. Khi gặp s += " world", Go sẽ không nối " world" vào cuối chuỗi "hello" hiện có.

3. Thay vào đó, Go sẽ cấp phát một vùng bộ nhớ mới đủ lớn để chứa chuỗi mới "hello world".

4. Nó sẽ sao chép nội dung của chuỗi cũ ("hello") và chuỗi mới (" world") vào vùng bộ nhớ mới này.

5. Cuối cùng, nó sẽ gán lại biến s để trỏ đến chuỗi mới.
```

Nếu bạn lặp lại thao tác này trong một vòng lặp, mỗi lần nối chuỗi sẽ tạo ra một chuỗi mới và sao chép dữ liệu, dẫn đến:

- ***Tiêu tốn bộ nhớ***: Nhiều chuỗi tạm thời được tạo ra, gây lãng phí.

- ***Gây áp lực lên Garbage Collector (GC)***: Trình dọn rác phải làm việc liên tục để thu dọn những chuỗi tạm thời không còn được sử dụng, làm giảm hiệu suất của chương trình.

***Cách giải quyết***
Giải pháp cho vấn đề này là sử dụng các kiểu dữ liệu có thể thay đổi `(mutable)` để xử lý dữ liệu, sau đó mới chuyển đổi về string khi cần thiết. Trong Go, cách tốt nhất là sử dụng `bytes.Buffer`.

`bytes.Buffer` là một cấu trúc dữ liệu được thiết kế để xây dựng các chuỗi một cách hiệu quả. Nó hoạt động như một "builder" (công cụ xây dựng):

- Nó quản lý một mảng `[]byte` có thể mở rộng.

- Khi bạn thêm dữ liệu vào `bytes.Buffer`, nó sẽ nối dữ liệu đó vào mảng byte nội bộ. Khi mảng này đầy, nó sẽ tự động cấp phát lại bộ nhớ, nhưng theo một cách hiệu quả hơn so với việc nối chuỗi thông thường.

- Sau khi hoàn thành, bạn chỉ cần gọi phương thức `.String()` để lấy chuỗi kết quả cuối cùng.

Ví dụ so sánh:

***Cách không hiệu quả (dùng string):***

```Go

// Nối chuỗi trong vòng lặp - KÉM HIỆU QUẢ
func inefficientConcat() string {
    s := ""
    for i := 0; i < 1000; i++ {
        s += "a"
    }
    return s
}
```

***Cách hiệu quả (dùng bytes.Buffer):***

```Go
// Dùng bytes.Buffer - RẤT HIỆU QUẢ
import (
    "bytes"
)

func efficientConcat() string {
    var b bytes.Buffer
    for i := 0; i < 1000; i++ {
        b.WriteString("a")
    }
    return b.String()
}
```
***Kết luận***: Với `bytes.Buffer`, bạn chỉ cần một lần cấp phát bộ nhớ cho chuỗi kết quả cuối cùng (khi gọi `.String()`), thay vì hàng nghìn lần cấp phát và sao chép như khi nối chuỗi trực tiếp. Điều này giúp mã của bạn nhanh hơn và tiết kiệm tài nguyên bộ nhớ hơn rất nhiều.

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
### Cấu trúc của slice trong Go
Một slice không phải là một mảng, mà là một cấu trúc dữ liệu gồm ***ba trường***: một con trỏ trỏ đến mảng underlying (mảng cơ sở), độ dài (length) và dung lượng (capacity).

- ***Length***: Số lượng phần tử hiện có trong slice.

- ***Capacity***: Số lượng phần tử tối đa mà mảng underlying có thể chứa.

Khi bạn sử dụng ***append*** để thêm một phần tử vào slice, Go sẽ kiểm tra xem *length* có bằng *capacity* hay không.

- Nếu *length* < *capacity*, Go chỉ đơn giản là thêm phần tử vào vị trí tiếp theo của mảng underlying. Việc này rất nhanh.

- Nếu *length* == *capacity*, lúc này mảng underlying đã đầy. Go sẽ phải làm một việc tốn kém:
```
1. Tạo một mảng underlying mới với dung lượng lớn hơn (thường là gấp đôi dung lượng cũ).

2. Sao chép tất cả các phần tử từ mảng cũ sang mảng mới.

3. Thêm phần tử mới vào mảng mới.

4. Cập nhật con trỏ của slice để trỏ đến mảng mới.
```

Đây chính là quá trình cấp phát lại bộ nhớ (***reallocation***). Việc này tốn kém về thời gian và tài nguyên, đặc biệt là khi slice chứa nhiều dữ liệu.

#### Mẹo: Sử dụng make để cấp phát trước
Việc ***preallocate*** (cấp phát trước) slice với *make* khi bạn có thể ***ước tính*** hoặc biết trước số lượng phần tử. Đây là cách giải quyết vấn đề cấp phát lại bộ nhớ một cách hiệu quả nhất.

Hãy cùng xem lại ví dụ trong tài liệu để thấy rõ sự khác biệt:

***1. Cách thông thường (có thể gây cấp phát lại)***

```
func BuildSlice(n int) []int {
    var s []int // Slice được khởi tạo rỗng, length = 0, capacity = 0
    for i := 0; i < n; i++ {
        s = append(s, i) // Mỗi lần append sẽ kiểm tra và có thể cấp phát lại
    }
    return s
}
```

Khi bạn chạy vòng lặp này, slice sẽ phải cấp phát lại nhiều lần. Ví dụ, khi `i=1`, capacity có thể tăng lên 2. Khi `i=2`, nó có thể tăng lên 4, rồi 8, 16... cứ như vậy. Mỗi lần cấp phát lại đều tốn chi phí.

***2. Cách dùng make để tối ưu (được khuyến khích)***
```Go
func BuildSlice(n int) []int {
    // Khởi tạo slice với length = 0 và capacity = n
    s := make([]int, 0, n)
    for i := 0; i < n; i++ {
        s = append(s, i) // Thao tác append sẽ không gây cấp phát lại cho đến khi đủ n phần tử
    }
    return s
}
```

Ở đây, chúng ta dùng make([]int, 0, n) để tạo một slice có:

- ***Length = 0***: Slice ban đầu trống, sẵn sàng để thêm phần tử.

- ***Capacity = n***: Dung lượng của mảng underlying ngay từ đầu đã đủ để chứa n phần tử.

Khi bạn chạy vòng lặp `for`, các thao tác `append` sẽ diễn ra cực kỳ nhanh chóng vì chúng không cần phải cấp phát lại bộ nhớ cho đến khi đạt đủ `n` phần tử. Điều này giúp loại bỏ hoàn toàn chi phí "ngầm" của việc cấp phát lại, làm cho code của bạn nhanh hơn và hiệu quả hơn về bộ nhớ, đặc biệt là với các slice lớn.


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

