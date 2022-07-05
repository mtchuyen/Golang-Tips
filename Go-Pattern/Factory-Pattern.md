# Factory Pattern

Các kỹ thuật sử dụng:
- Structs & Interfaces.

***Sức mạnh*** của Factory Pattern là:
>> ...create objects without having to specify the exact 'type' of the object that will be created...
---> chính là cách sử dụng ***interface*** --> Đọc trong [3]


Các chú ý nhắc lại:
- Factory pattern là một trong 9 Pattern phổ biến trong lập trình hướng đối tượng (OOP). 
- Factory pattern thuộc "Nhóm khởi tạo" (Creational Pattern) trong 23 pattern.
- Nên đọc `Builder Pattern` trước.

Rất hay được sử dụng vì: writing clean, concise, maintainable & testable code.

## Factory pattern
- Factory pattern cung cấp *cách tốt nhất* để tạo ra các đối tượng.
- Nó giúp cho chúng ta định nghĩa được nhiều đối tượng và cho phép các lớp con tự quyết định là cái nào được khởi tạo.

Có 3 loại mô hình đó là: 
- Simple factory 
- Factory method: Factory method là một pattern nằm trong nhóm Polymorphic Factory (đa hình). Và vì tính trừu tượng của nó, Factory Method còn được gọi là Virtual Constructor. 
- Abstract factory

## A Class Factory in Golang (Google Go)

http://www.minaandrawos.com/2014/09/25/class-factory-in-golang-google-go/

Chúng ta xem xét 1 interface có tên là `AbstractFactory`:

```go
type AbstractFactory interface {
        CreateMyLove()
}
```
Ở đây có 2 khái niệm:
- Interface `AbstractFactory` có method là `CreateMyLove()`
- `CreateMyLove()` được gọi là 1 ***factory method***.


## Tham khảo:
[1] https://techmaster.vn/posts/36708/factory-pattern-trong-golang

[2] https://anonystick.com/blog-developer/phan-2-factory-pattern-cach-ma-toi-trien-khai-trong-nha-may-vinfast-fresher-va-junior-nen-bo-qua-phan-3-2020110554662242

[3] https://matthewbrown.io/2016/01/23/factory-pattern-in-golang/

#### Series bài viết `Factory Pattern` của Pravand Katyare:
[U.1] https://medium.com/@katyare.pravand/factory-design-pattern-in-go-63a9b423976f

[U.2] https://towardsdev.com/interface-factory-design-pattern-in-go-3eb6f97f20c4

[U.3] https://medium.com/@katyare.pravand/factory-generator-pattern-in-go-ca6e9af12a1b
