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

## Design Principles có liên quan gì đến Design Patterns không ???
Câu trả lời là có, nhưng có chút khác biệt.
- ***Principles***: là các hướng dẫn, mang tính trừu tượng nhiều hơn.
- ***Patterns***: là những ví dụ cụ thể hóa, cung cấp các giải pháp tái sử dụng đến các vấn đề thực tế. Patterns tốt nên tuân thủ tốt Principles.

## Các nhóm Design Principles
Trong cuốn "Agile Software Development: Principles, Patterns, and Practices", Uncle Bob đã tổng hợp lại 11 nguyên lý và chia chúng làm 3 nhóm:

***Class Design principles – bao gồm 5 nguyên lý***

- (SRP) Single Responsibility Principle
- (OCP) Open Closed Principle
- (LSP) Liskov Substitution Principle
- (DIP) Dependency Inversion Principle
- (ISP) Interface Segregation Principle

> viết tắt là S.O.L.I.D (hoặc SOLID)

***Package Cohesion Principles - bao gổm 3 nguyên lý***

- (REP) Reuse-Release Equivalence Principle
- (CRP) Common Reuse Principle
- (CCP) Common Closure Principle


***Package Coupling principle - bao gồm 3 nguyên lý***
- (ADP) Acyclic Dependencies Principle
- (SDP) Stable Dependencies Principle
- (SAP) Stable Abstractions Principle


***Ngoài ra còn có các nguyên lý khác như***
- (DRY) Don't Repeat Yourself
- (KISS) Keep It Simple Stupid ...

## Nguyên tắc KISS 
Source: (Ref.8)

***KISS = Keep It Simple Stupid***

KISS có nhiều biến thể khác nhau như "Keep It Short and Simple", "Keep It Simple and Straightforward" và "Keep It Small and Simple".

Tựu chung lại, hàm ý của nó vẫn hướng về một sự đơn giản và rõ ràng trong mọi vấn đề. Và như vậy, sự đơn giản là mục đích trọng tâm trong thiết kế, còn những cái phức tạp không cần thiết thì nên tránh.

Trong lập trình, KISS nghĩa là hãy làm cho mọi thứ (mã lệnh của bạn) trở nên đơn giản và dễ nhìn hơn. Hãy chia nhỏ vấn đề ra những bài toán nhỏ đơn giản hơn, viết mã để xử lý nó, biến nó thành các lớp và các hàm riêng biệt, đừng để một lớp hay một phương thức có hàng trăm dòng lệnh, hãy để nó trở về con số chục thôi.

Đừng viết những lớp hay phương thức theo kiểu spaghetti hay all-in-one (tất cả trong một) như dao thụy sĩ, hãy để mọi thứ thật đơn giản để bạn luôn có thể hiểu được, và kết hợp chúng với nhau để giải quyết được các bài toán lớn.

## Nguyên tắc YAGNI:
Source: (Ref.8)

***YAGNI = You Aren’t Gonna Need It***

Là những lập trình viên, đôi khi chúng ta suy nghĩ quá nhiều về tương lai của dự án nên chúng ta code thêm thật nhiều tính năng “phòng khi chúng ta cần đến nó” hay “Cuối cùng chúng ta sẽ dùng đến nó”. Ý nghĩ này sai hoàn toàn. Bạn đã không cần đến nó thì bạn sẽ không cần đến nó trong hầu hết tất cả các trường hợp! “You Aren’t Gonna Need It!”.

Nếu bạn nghĩ một chức năng sẽ hữu dụng trong tương lai, hãy cứ bình tĩnh và xem lại những công việc đang chờ bạn giải quyết ngay. Bạn không thể lãng phí thời gian của mình để code những tính năng mà bạn sẽ phải sửa nó hay thay đổi nó bởi vì nó không phù hợp với nhu cầu của khách hàng, hay trong trường hợp xấu nhất là nó sẽ không được sử dụng đến.

> tốt nhất là hãy giải quyết thành công vấn đề hiển hiện trước mắt đã, khách hàng không vẽ việc cho bạn thì bạn cũng đừng tự vẽ việc cho mình

## Nguyên tắc DRY:
Source: (Ref.8)

***DRY = Don’t Repeat Yourself***

Nguyên tắc DRY chỉ ra rằng nếu chúng ta đang muốn viết nhiều đoạn code giống nhau ở nhiều chỗ khác nhau, thay vì copy và paste đoạn code đó, chúng ta hãy đưa đoạn code đó vào một phương thức riêng sau đó gọi phương thức này từ những chỗ chúng ta cần gọi.

Quen quen đúng không, vì nó giống như tính chất kế thừa trong lập trình hướng đối tượng OOP mà chúng ta đã quá quen thuộc rồi 

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
- [8:Những nguyên tắc, định luật thông dụng khi lập trình](https://viblo.asia/p/nhung-nguyen-tac-dinh-luat-thong-dung-khi-lap-trinh-gGJ59g2DZX2)

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
- 
