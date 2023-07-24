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


### 1. Method declaration
- Define: https://golang.org/ref/spec#Method_declarations
- [Method declaration with function receiver in Golang](https://pgillich.medium.com/method-declaration-with-function-receiver-in-golang-7f5531ded97d)

### 2. Method theory
The term "method" came up with object-oriented programming. 
> In an OOP language (like C++ for example) you can define a "class" which encapsulates data and functions which belong together.
> Those functions inside a class are called "methods" and you need an instance of that class to call such a method.
See: [whats-the-difference-of-functions-and-methods-in-go](https://stackoverflow.com/questions/8263546/whats-the-difference-of-functions-and-methods-in-go)

### Ref:
- [Ref2: whats-the-difference-of-functions-and-methods-in-go](https://stackoverflow.com/questions/8263546/whats-the-difference-of-functions-and-methods-in-go)
- [Ref3: Golang functions vs methods](https://www.sohamkamani.com/golang/functions-vs-methods/)
- [Ref4: Series 1: Golang function vs method Comparison with Examples](https://www.golinuxcloud.com/golang-function-vs-method/)
- [Ref5: Series 2: GO Methods vs Functions [In-Depth Explanation]](https://www.golinuxcloud.com/go-methods-vs-functions/)
