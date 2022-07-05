***Gang of Four*** viết ra 23 `design pattern` nhằm giúp các developer có hướng đi rõ ràng để thiết kế phần mềm hoàn chỉnh. 23 design patterns chia làm 3 patterns lớn là:
- Creation
- Structure
- Behavior

### 9 mẫu thiết kế mà mỗi lập trình viên dều phải biết

***phần 1: Builder pattern***

https://anonystick.com/blog-developer/9-mau-thiet-ke-ma-moi-lap-trinh-vien-deu-phai-biet-phan-1-builder-pattern-2020103194916615

***Phần 2: Factory pattern***

https://anonystick.com/blog-developer/phan-2-factory-pattern-cach-ma-toi-trien-khai-trong-nha-may-vinfast-fresher-va-junior-nen-bo-qua-phan-3-2020110554662242

### golang-design-pattern

- https://github.com/anhthii/golang-design-pattern
- https://refactoring.guru/design-patterns/go

### Builder Pattern trong golang

- https://techmaster.vn/posts/36713/builder-pattern-trong-golang


### [Factory Pattern trong golang](https://github.com/mtchuyen/Golang-Tips/blob/master/Go-Pattern/Factory-Pattern.md)


### Singleton
- https://viblo.asia/p/go-design-pattern-singleton-3P0lPeYm5ox
- https://locphamtan.medium.com/singleton-pattern-trong-golang-dbe81b93b41a


### Adapter Design Pattern

- https://viblo.asia/p/adapter-pattern-n7prv3mOMKod
- https://gpcoder.com/4483-huong-dan-java-design-pattern-adapter/
- https://devhoi.com/threads/adapter-pattern-la-gi-series-design-pattern.33/
- https://levelup.gitconnected.com/adapter-design-pattern-in-golang-template-example-5fcd2165c44

### options pattern in golang

See : `Options-Pattern.md`

https://levelup.gitconnected.com/options-pattern-in-golang-9a0384a9d8db

### Observer Design Pattern
***Định nghĩa 1***: ***Observer*** là một `behavioral design pattern` (sau đây gọi tắt là DP) dùng để định nghĩa một cơ chế đăng ký (subscribe) nhằm thông báo (notify) cho nhiều đối tượng về các sự kiện xảy ra với đối tượng mà chúng đang quan sát (observe).

***Định nghĩa 2***: ***Observer*** pattern là một mẫu thiết kế phần mềm mà một đối tượng - gọi là ***subject***, duy trì một danh sách các thành phần phụ thuộc nó - gọi là ***observer***, và thông báo tới chúng một cách tự động về bất cứ thay đổi nào, thường thì bằng cách gọi 1 hàm của chúng.

***đối tượng*** - gọi là ***subject***: là một `thứ` mà có một/nhiều trạng thái mà người/vật/thứ-khác quan tâm tới.
- VD: `iPhone` là một `thứ`, `mới ra` là trạng thái. Nên `iPhone mới ra` là trạng thái của iPhone được `giới-công-nghệ` quan tâm.
- 
#### Bài toàn:
Tưởng tượng rằng bạn có hai loại đối tượng: `Khách hàng` và `Cửa hàng`. Khách hàng rất quan tâm đến một thương hiệu sản phẩm cụ thể (giả sử là một mẫu iPhone mới) chuẩn bị được bày bán tại cửa hàng.

Khách hàng có thể ghé thăm cửa hàng mỗi ngày để kiểm tra đã có sản phẩm chưa. Nhưng trong khi sản phẩm vẫn chưa được tung ra, hầu hết những chuyến đi tới cửa hàng này sẽ là vô nghĩa.

Mặt khác, cửa hàng có thể gửi hàng tấn email (mà sẽ có khách hàng nghĩ là thư rác) cho `tất cả khách hàng` (kể cả các `khách hàng` không quan tâm tới iPhone mới) mỗi khi có sản phẩm mới. Điều này sẽ giúp một số khách hàng đỡ phải mất công đến cửa hàng vô số lần, tuy nhiên, đồng thời, điều này sẽ làm khó chịu những khách hàng khác không quan tâm đến sản phẩm mới.

Có vẻ như có một sự xung đột xảy ra ở đây. Hoặc là khách hàng lãng phí thời gian kiểm tra xem sản phẩm đã có hàng hay chưa, hoặc là cửa hàng lãng phí nguồn lực khi thông báo cho khách hàng không muốn nhận thông báo.

#### Mục tiêu
- Định nghĩa mối phụ thuộc `một - nhiều` (1-n) giữa các đối tượng để khi mà một đối tượng có sự thay đổi trạng thái, tất các thành phần phụ thuộc của nó sẽ được thông báo và cập nhật một cách tự động.
- Một đối tượng có thể thông báo đến một số lượng không giới hạn các đối tượng khác.

***Ví dụ:***

Ta có 1 bảng số liệu, từ bảng đó ta có thể vẽ được `n` biểu đồ.
- Như vậy ta có mối quan hệ `1-n`.
- Bảng số liệu gọi là `subject`, biểu đồ gọi là `observer`: các `biểu đồ` (n) phụ thuộc vào `bảng số liệu`.
- Khi ta thay đổi giá trị của 1 ô (cột, hàng) của bảng số liệu, thì các biểu đồ sẽ thay đổi theo: `biểu đồ` phụ thuộc vào `bảng số liệu`

#### Phạm vi ứng dụng
Observer Pattern được áp dụng khi:
- Sự thay đổi trạng thái ở 1 đối tượng có thể được thông báo đến các đối tượng khác mà không phải giữ chúng liên kết quá chặt chẽ

#### Tham khảo:
- [Observer Design Pattern in Golang with an Example](https://levelup.gitconnected.com/observer-design-pattern-in-golang-with-an-example-6c24898059b1)
- https://viblo.asia/p/gioi-thieu-ve-observer-design-pattern-Ljy5V4j9Zra
- https://viblo.asia/p/design-pattern-observer-V3m5WO8W5O7
- https://codelearn.io/sharing/observer-pattern-va-ung-dung
- https://viblo.asia/p/gioi-thieu-ve-observer-design-pattern-Ljy5V4j9Zra
- https://topdev.vn/blog/observer-pattern-la-gi-nhat-tru-kinh-thien/
- [rất dài]: https://toihocdesignpattern.com/head-first-design-patterns-tieng-viet-chuong-2-observer-pattern.html

## Clean architecture

[Golang and clean architecture](https://itnext.io/golang-and-clean-architecture-19ae9aae5683)

