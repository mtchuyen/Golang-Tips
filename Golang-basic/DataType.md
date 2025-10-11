# 1. Tổng quan phân loại Data Types trong Go

| Nhóm chính          | Mô tả                                                     | Đặc điểm nổi bật                                     |
| ------------------- | --------------------------------------------------------- | ---------------------------------------------------- |
| **Basic types**     | Các kiểu dữ liệu cơ bản, nguyên thủy                      | Giá trị độc lập, lưu trực tiếp                       |
| **Aggregate types** | Gom nhiều giá trị cơ bản thành 1 đơn vị                   | Gồm `array` và `struct`                              |
| **Reference types** | Kiểu dữ liệu tham chiếu (trỏ tới vùng nhớ khác)           | Gồm `slice`, `map`, `channel`, `function`, `pointer` |
| **Interface types** | Định nghĩa hành vi (method set), không lưu giá trị cụ thể | Cho phép polymorphism và abstraction                 |

# 2. Basic Types (Kiểu cơ bản)

| Nhóm con                    | Các kiểu                                                 |
| --------------------------- | -------------------------------------------------------- |
| **Số nguyên**               | `int`, `int8`, `int16`, `int32`, `int64`                 |
| **Số nguyên không dấu**     | `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` |
| **Số thực**                 | `float32`, `float64`                                     |
| **Số phức**                 | `complex64`, `complex128`                                |
| **Ký tự / Chuỗi**           | `byte`, `rune`, `string`, `bool`                         |
| **Boolean**                 | `bool`                                                   |

Ex:
```go
var a int = 10
var b float64 = 3.14
var c complex128 = complex(2, 3)
var s string = "Hello"
var ok bool = true
```
***Note:***
- `string` là ***immutable*** (không thể thay đổi nội dung sau khi tạo).
- `rune` là alias của `int32` (đại diện cho 1 ký tự Unicode).
- `byte` là alias của `uint8`.

## 2.1. Kiểu bool

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
## 2.2. Kiểu dữ liệu số (numeric type)
### 2.2.1. Số nguyên
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

### 2.2.2. Số thực
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

### 2.2.3. Số phức
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


### 2.2.4. Các kiểu dữ liệu số khác

- **byte** là tên gọi khác của **uint8**
- **rune** là tên gọi khác của **int32**

---
### 2.3. Kiểu dữ kiệu chuỗi

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


# 3. Aggregate Types (Kiểu tập hợp)
## 3.1. Array
## 3.2. Struct

# 4. Reference Types (Kiểu tham chiếu)
Các biến thuộc nhóm này ***không*** lưu trực tiếp giá trị, mà lưu địa chỉ tham chiếu đến vùng nhớ chứa dữ liệu thực.

==> Vì vậy, khi gán hoặc truyền vào hàm, chúng ***chia sẻ cùng một dữ liệu gốc (không copy).***

## 4.1. Slice
## 4.2. Map
## 4.3. Channel
## 4.4. Pointer
## 4.5. Function

# 5. Interface Types
- Interface mô tả ***bộ hành vi*** (method set) mà một kiểu cần thực hiện.
- Không chứa dữ liệu cụ thể, mà chứa ***(type, value)*** của biến thực thi.

Ex:
```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
    fmt.Println(s.Area())
}

func main() {
    c := Circle{Radius: 2}
    PrintArea(c)
}
```

# Tổng hợp

| Nhóm          | Ví dụ                                     | Lưu trữ                      | Truyền vào hàm            | Ghi chú                  |
| ------------- | ----------------------------------------- | ---------------------------- | ------------------------- | ------------------------ |
| **Basic**     | `int`, `string`, `bool`                   | Giá trị trực tiếp            | Copy toàn bộ              | Kiểu nguyên thủy         |
| **Aggregate** | `array`, `struct`                         | Tập hợp nhiều giá trị        | Copy toàn bộ              | Có thể chứa nested types |
| **Reference** | `slice`, `map`, `chan`, `func`, `pointer` | Tham chiếu tới vùng nhớ khác | Cùng chia sẻ vùng dữ liệu | Không copy giá trị thật  |
| **Interface** | `interface{}`, `io.Reader`                | (type, value)                | Theo hợp đồng hành vi     | Hỗ trợ trừu tượng hóa    |

```pgsql

Data Types in Go
├── Basic
│   ├── Numbers (int, float, complex)
│   ├── Boolean
│   └── String (immutable)
│
├── Aggregate
│   ├── Array (fixed size)
│   └── Struct (custom type)
│
├── Reference
│   ├── Slice
│   ├── Map
│   ├── Channel
│   ├── Function
│   └── Pointer
│
└── Interface
    ├── Defines method set
    ├── Enables polymorphism
    └── Example: io.Reader, fmt.Stringer
```
