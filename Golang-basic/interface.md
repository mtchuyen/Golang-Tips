## interface

***Note***: trước khi đọc `Interface` cần nằm được `struct` (structure) và `method` (function của struct) ở [phần đầu](https://github.com/mtchuyen/Golang-Tips/tree/master/Golang-basic#method)

### Định nghĩa về Interface

Định nghĩa chung về ***interface*** trong thế giới `Hướng đối tượng` đó là `Interface xác định CÁC hành vi của một đối tượng`. Nó chỉ xác định những gì đối tượng phải làm. 

An ***interface*** is a collection of ***method signatures*** that a ***Type*** can implement (using methods). So interface defines (*not declares*) the behavior of the object (of the type `Type`).
- `Method signatures`: Thông tin cần thiết để gọi method.
- `defines`: Interface sẽ xác định các hành vi.
- `declares`: quá trình hoạt động của 1 hành vi.

***Ví dụ***: Trích từ [1-medium]

```go
type Shape interface {
	Area() float64
	Perimeter() float64
}
func main() {
	var newS Shape 
}
```
Giải thích:
- `Shape` là một interface, có 2 method: `Area` và `Perimeter`
- `Shape` được khai báo (define) với 2 hành vi (method), nhưng không mô tả rõ hành động (declare)
- `Shape` là một kiểu dữ liệu (Type), được khai báo như các kiểu dữ liệu thông thường khác (string, int,...)

### Nguyên lý/cách thực thi (implement) của Interface

`Interface trong Go` tuân theo khái niệm về `concept` [Duck test](https://en.wikipedia.org/wiki/Duck_test) như sau:

```If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.```

Nghĩa là : ```nếu con nào trông giống con vịt, bơi như vịt, quạc quạc như vịt thì con đó đích thực là con vịt```.

Đọc thêm về bài viết ***Nhận dạng vịt***:
- [wikipedia](https://vi.wikipedia.org/wiki/Nh%E1%BA%ADn_d%E1%BA%A1ng_v%E1%BB%8Bt)

Qua đó ta có 1 cách hiểu dữ liệu gọi là ***Duck type***:
- Nếu 1 con vật mà ta thấy nó giống như vịt, đi như vịt, kêu như vịt, thì đó phải là con vịt.
- Nếu con chó bị chủ trang điểm thành có hình dáng giống vịt, chó lại biết kêu như vịt, đi như vịt. Thì con chó này vẫn bị gọi là vịt.

Qua ví dụ trên, ta thấy rằng `con chó` mà có các hành vi như vịt, sẽ hiểu như struct có các hành vi của interface. Ta gọi `con chó` đang thực thi (implement) nhiệm vụ của vịt, hay `type` đang thực thi (impletment) của interface.

Vậy nên quy định trong Golang là: nếu có một kiểu dữ liệu chứa tất cả các phương thức được khai báo trong interface đó, thì lúc này Interface đã được implement.

***Ví dụ:*** trích từ [3:medium.com]

```go
package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectange s", s.Area())
	fmt.Println("s == r is", s == r)
}

```

Trong ví dụ trên, ta có ***interface*** `Shape` và ***struct*** `Rect`. Sau đó ta define các ***method*** `Area`, `Perimeter` thuộc về receiver `Rect`. Vì vậy, Rect implement các method Area, Perimeter.

Vì 2 method này được define trong interface `Shape`, nên struct Rect ***implement*** interface Shape, việc này xảy ra tự động.

Điều này có nghĩa là gì:
- `s` có type là `Shape`, `r` có type là `Rect`, nhưng vì `Rect` thực thi `Shape` nên lúc này `s==r` và kết quả của dòng cuối cùng sẽ in ra là: `s == r is true`.
- việc này xảy ra tự động, nên khi lập trình chúng ta không cần check kiểu (type) và ép kiểu.



***Tại sao lại phải quy định như thế?***

1. Để cho tiện khai báo, gán dữ liệu:

Ví dụ: Xem phần `[Công dụng thực tế của interface]` ở phía dưới, ta thấy các đối tượng pemp1, cemp1 có `type` khác nhau, nhưng đều chung 1 interface `SalaryCalculator`, nên có thể cho chung vào 1 slice được.

### Công dụng thực tế của interface
- nguồn: [1:techmaster.vn] (note: đưa vào đây vì sợ rằng có 1 ngày techmaster.vn change link/domain)

Thực tế là 1 interface được `define`, nhưng các hoạt động lại cần thiết ở `declare`, mà `declare` lại được khai báo ở `struct` và `method`. Vậy câu hỏi đặt ra là tại sao lại không sử dụng luôn `struct` và `method`?

Ở ví dụ sau đây, chúng ta sẽ thấy tác dụng của `interface`.

Chúng ta viết một chương trình đơn giản tính tổng chi phí của một công ty dựa trên mức lương của từng nhân viên. Giả định tất cả các chi phí tính bằng USD.


```go
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  //nhân viên chính thức
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  //nhân viên hợp đồng
    empId  int
    basicpay int
}

//tiền lương của nhân viên permanent bằng tổng của basic pay và pf
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//tiền lương của nhân viên contract chỉ là basic pay
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

//Có 2 method có khai báo tên giống nhau: CalculateSalary()

/*
tổng chi phí được tính bằng cách duyệt qua từng phần tử của slice SalaryCalculator
và tính tổng mức lương của từng nhân viên
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

}
```
Ta thấy:
- Có 2 method có khai báo tên giống nhau: `CalculateSalary()`, nhưng lại của 2 `struct` khác nhau.
- hai báo một `interface` tên là `SalaryCalculator` có một phương thức là `CalculateSalary()` return `int`.

Chúng ta có 2 loại nhân viên trong công ty là `Permanent` (nhân viên chính thức) và `Contract` (nhân viên hợp đồng) được định nghĩa bằng kiểu `struct`.
- Mức lương của nhân viên `Permanent` là tổng của `basicpay` và `pf` còn đối với nhân viên Contract thì chỉ là basicpay. 
- Điều này được thể hiện trong các phương thức `CalculateSalary` tương ứng. Bằng cách khai báo phương thức này, cả 2 struct `Permanent` và `Contract` đều đang ***implement*** interface `SalaryCalculator`.

Hàm `totalExpense` được khai báo bên dưới ***thể hiện sự tiện ích của việc sử dụng interface.***: Hàm này nhận một slice các interface `SalaryCalculator []SalaryCalculator` làm tham số. 

Trong hàm main() chúng ta truyền một slice với các phần tử gồm cả 2 kiểu `Permanent` và `Contract` vào hàm `totalExpense`. Hàm totalExpense tính toán chi phí bằng cách gọi đến phương thức (method) `CalculateSalary` của kiểu tương ứng, điều này được thực hiện ở câu lệnh expense = expense + v.CalculateSalary().

Ưu điểm lớn nhất của hàm totalExpense này là nó *có thể được mở rộng đến bất kỳ loại nhân viên mới nào* mà không cần phải thay đổi code. Giả sử công ty bổ sung một loại nhân viên mới là `Freelancer` với cách tính lương khác. Freelancer này chỉ việc truyền vào đối số slice của hàm totalExpense mà không phải thay đổi bất kỳ 1 dòng code nào trong hàm totalExpense. Freelancer cũng implement interface SalaryCalculator.

Output của chương trình trên là: Total Expense Per Month $14050

### Type: đại diện cho kiểu giá trị mà chúng ta sử dụng. 

***Tác dụng:***

- Nhờ có `Type` mà các compiler có thể xác định được một số lỗi trong quá trình compile.

***Type assertion:***

- assertion: sự xác nhận
- là một cách để thông báo cho compiler biết kiểu của đối tượng mà mình sử dụng.
- Nói ngắn gọn là ép kiểu cho đối tượng.


### Type assertions

***Type assertion*** cung cấp tính năng kiểm tra kiểu dữ liệu.

```
t := i.(T)
```
*Giải thích syntax:*

- Kiểm tra `i` có thuộc kiểu dữ liệu `T` hay không. Nếu đúng thì gán `t = i`. Nếu không sẽ trigger một panic (lỗi).

- Để test giá trị nhưng ***không*** trigger panic, ta dùng syntax sau:

```
t, ok := i.(T)
```
*Giải thích:*
- Nếu `i` có thuộc kiểu dữ liệu `T` thì gán `t = i`, `ok = true`. Nếu không, `t` sẽ là `zero` value, `ok = false`


## Ref:

[1] https://techmaster.vn/posts/35059/series-golang-co-ban-phan-18-interfaces-i

[2] https://backend.vn/bai-viet/golang-interface-1541487108.html

[3] https://medium.com/@anh.nt094/golang-eb65bfe1a8bb

[4] https://research.swtch.com/interfaces

