# Introduce
Phần này sẽ tập trung vào ***Object Oriented Design*** 
(Ref.4)Ngày xửa ngày xưa, người ta nói rằng:

> Proper Object Oriented Design makes a developer's life easy, whereas bad design makes it a disaster

Vâng good design sẽ làm cho cuộc sống của lập trình viên dễ dàng hơn. Vậy làm sao để có một good design ? Chúng ta có những nguyên lý (principles).
> Principles: Là một tập hợp các hướng dẫn giúp chúng ta tránh các thiết kế xấu.


## Tại sao cần tuân thủ 1 nguyên lý trong lập trình:

Tuân thủ nguyên lý (principles) giống như cơ thể có 1 bộ khung xương (nguyên lý tốt thì khung xương đẹp vững chắc,...) vậy nên cần chọn 1 nguyên lý phù hợp với dự án.
> (Ref.4)Việc viết code theo các nguyên lý giúp cho code trở nên trong sáng, dễ đọc, dễ test và quan trọng nhất là dễ maintain hơn rất nhiều. Và chúng ta đều biết, trong vòng đời của một phần mềm, thời gian code chỉ chiếm 20-40% còn lại là thời gian dành cho maintain: thêm/bớt chức năng, fix bug...

Chung quy lại vẫn là để KHÔNG bị ***Bad code***.

(Ref.1) Theo Robert C. Martin, một trong những đồng tác giả của Agile Manifesto, có ba đặc điểm của thiết kế xấu cần tránh:
- ***Rigidity (cứng nhắc)***: Đề cập đến là việc khó khăn khi sửa đổi ứng dụng vì mọi thay đổi liên quan đến quá nhiều phần của hệ thống.
- ***Fragility (phá vỡ)***: Đây là sự phát sinh lỗi trong ứng dụng do những thay đổi trong các phần khác.
- ***Immobility (bất động)***: Đây là việc hệ thống không thể sử dụng thêm một thành phần trong phần mềm khác vì nó quá phụ thuộc vào ứng dụng hiện tại.

(MTC)Nhưng trong (Ref.2) phần ***Bad code*** liệt kê là:
- Rigid. Is the code rigid? Does it have a straight jacket of overbearing types and parameters, that making modification difficult?
- Fragile. Is the code fragile? Does the slightest change ripple through the code base causing untold havoc?
- Immobile. Is the code hard to refactor? Is it one keystroke away from an import loop?
- ***Complex***. Is there code for the sake of having code, are things over-engineered?
- ***Verbose***. Is it just exhausting to use the code? When you look at it, can you even tell what this code is trying to do?


# SOLID

(Ref.1) SOLID là một từ viết tắt thường dùng để chỉ một bộ năm nguyên tắc cơ bản của một thiết kế mô hình:
- ***S***: Single Responsibility Principle
- ***O***: Open/Closed Principle
- ***L***: Liskov Substitution Principle
- ***I***: Interface Segregation Principle
- ***D***: Dependency Inversion Principle

Để hiểu về 5 nguyên tắc của SOLID, trước hết xem qua giải thích bằng hình ảnh ở (Ref.3).

## Tác dụng của SOLID
- Các nguyên tắc SOLID có thể giúp chúng ta thiết kế các ứng dụng tốt hơn (Ref.1). 
- Chúng cho phép chúng ta phát hiện các điểm yếu và có nhưng phần mềm với dòng code mạnh mẽ, linh hoạt và đặc biệt là khả năng bảo trì. Khả năng bảo trì và tính linh hoạt của một ứng dụng chủ yếu phụ thuộc vào cách thiết kế, cách kết hợp các thành phần và sử dụng các nguyên tắc (Ref.1).

## 5 nguyên tắc của SOLID

## 1. Single Responsibility Principle
Nguyên tắc đầu tiên của ngăn xếp SOLID là Nguyên tắc Single Responsibility.

## Ref:
- [1:Nguyên lý SOLID là gì...](https://viblo.asia/p/nguyen-ly-solid-la-gi-dependency-injection-la-gi-bWrZnX4b5xw)
- [2: SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)
- [3: Phương pháp thiết kế hướng đối tượng S.O.L.I.D qua hình ảnh](https://viblo.asia/p/phuong-phap-thiet-ke-huong-doi-tuong-solid-qua-hinh-anh-WAyK82jelxX)
- [4:Sơ lược Object Oriented Design Principles](https://viblo.asia/p/so-luoc-object-oriented-design-principles-MdZGAQODGox)

# SOLID in Golang

(Ref.6) nói chi tiết về nguyên lý SOLID với các ví dụ bằng ngôn ngữ Go.

## 1. Single Responsibility Principle

(Ref.6) nói chi tiết về nguyên tắc ***SRP***


## Ref: 
- [5: Writing SOLID Go code](https://bitek.dev/blog/go_solid/)
- [6:Tập tành SOLID với Golang: Single Responsibility Principle](https://viblo.asia/p/tap-tanh-solid-voi-golang-single-responsibility-principle-4P856GbLKY3)

## Code tham khảo
- [10: SOLID principles in Golang](https://github.com/ammorteza/SOLID-principles-in-Golang)
- [11: SOLID Principle in Golang](https://github.com/mcholismalik/solid-in-golang)
