
# SOLID

(Ref.1) SOLID là một từ viết tắt thường dùng để chỉ một bộ năm nguyên tắc cơ bản của một thiết kế mô hình:
- ***S***: Single Responsibility Principle
- ***O***: Open/Closed Principle
- ***L***: Liskov Substitution Principle
- ***I***: Interface Segregation Principle
- ***D***: Dependency Inversion Principle

Để hiểu về 5 nguyên tắc của SOLID, trước hết xem qua giải thích bằng hình ảnh ở (Ref.3) và (Ref.4).

## Tác dụng của SOLID
- Các nguyên tắc SOLID có thể giúp chúng ta thiết kế các ứng dụng tốt hơn (Ref.1). 
- Chúng cho phép chúng ta phát hiện các điểm yếu và có nhưng phần mềm với dòng code mạnh mẽ, linh hoạt và đặc biệt là khả năng bảo trì. Khả năng bảo trì và tính linh hoạt của một ứng dụng chủ yếu phụ thuộc vào cách thiết kế, cách kết hợp các thành phần và sử dụng các nguyên tắc (Ref.1).

## 5 nguyên tắc của SOLID

## 1. Single Responsibility Principle
Nguyên tắc đầu tiên của ngăn xếp SOLID là Nguyên tắc Single Responsibility.

## 2.  Open/Closed Principle

Viết tắt: ***OCP***

(Ref.6)*Robert C. Martin* đã coi nguyên tắc này là nguyên tắc quan trọng nhất của thiết kế hướng đối tượng. Nhưng ông ấy không phải là người đầu tiên định nghĩa nó. Bertrand Meyer đã viết về nó vào năm 1988 trong cuốn sách Object-Oriented Software Construction. Ông giải thích Nguyên tắcOpen/Closed Principle là:

***Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification***

Nguyên tắc này có thể được hiểu rằng: các class hoặc hàm nên được thiết kể để mở rộng, nhưng đóng lại để *tránh sự thay đổi trực tiếp* mã nguồn. Điều này có nghĩa là hệ thống nên được thiết để hỗ trợ các lập trình viên sau này có extend các class có sẵn để cung cấp thêm chức năng thay vì chỉnh sửa trên mã nguồn tồn tại sẵn trong hệ thống.

Đọc thêm:
- https://viblo.asia/p/lam-chu-solid-open-closed-principle-phan-2-djeZ1Gom5Wz


## Ref:
- [1:Nguyên lý SOLID là gì...](https://viblo.asia/p/nguyen-ly-solid-la-gi-dependency-injection-la-gi-bWrZnX4b5xw)
- [2:SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)
- [3:Phương pháp thiết kế hướng đối tượng S.O.L.I.D qua hình ảnh](https://viblo.asia/p/phuong-phap-thiet-ke-huong-doi-tuong-solid-qua-hinh-anh-WAyK82jelxX)
- [4:Sơ lược Object Oriented Design Principles](https://viblo.asia/p/so-luoc-object-oriented-design-principles-MdZGAQODGox)
- [7:Nguyên tắc thiết kế SOLID - Open/Closed Principle](https://viblo.asia/p/nguyen-tac-thiet-ke-solid-openclosed-principle-eW65GezJZDO)


# SOLID in Golang

(Ref.5) nói chi tiết về nguyên lý SOLID với các ví dụ bằng ngôn ngữ Go.

## 1. Single Responsibility Principle

(Ref.6) nói chi tiết về nguyên tắc ***SRP***

## 2. Open/Closed principle
Phần [Open/Closed principle](https://viblo.asia/p/lap-trinh-huong-doi-tuong-trong-phpphan-3-Az45bb1O5xY) có hình ảnh mô tả cho dễ hiểu.

## Ref: 
- [5: Writing SOLID Go code](https://bitek.dev/blog/go_solid/)
- [6:Tập tành SOLID với Golang: Single Responsibility Principle](https://viblo.asia/p/tap-tanh-solid-voi-golang-single-responsibility-principle-4P856GbLKY3)

## Code tham khảo
- [10: SOLID principles in Golang](https://github.com/ammorteza/SOLID-principles-in-Golang)
- [11: SOLID Principle in Golang](https://github.com/mcholismalik/solid-in-golang)
- [12:Practical SOLID in Golang: Interface Segregation Principle](https://levelup.gitconnected.com/practical-solid-in-golang-interface-segregation-principle-f272c2a9a270)
- [13:SOLID principle in GO](https://s8sg.medium.com/solid-principle-in-go-e1a624290346)
- [14:SOLID Principles in Golang Explained with Examples](https://levelup.gitconnected.com/solid-principles-in-golang-explained-by-examples-4a4cccf47388)
- [15:SOLID principles in Golang](http://thepedestrian.in/posts/2017-03-23-solid-principals/)
- [16:The advantages of go language in solid principle](https://developpaper.com/the-advantages-of-go-language-in-solid-principle/)
- 


## Ref
- [20](https://viblo.asia/p/cac-phuong-phap-thiet-ke-huong-doi-tuong-solid-DZrGNDgrkVB)
- [21](https://viblo.asia/p/dependency-inversion-inversion-of-control-and-dependency-injection-qzakzNYBkyO)
- [22](https://medium.com/@lucapelosi/5-design-principles-from-a-software-architecture-guru-8762a304fb3b)


