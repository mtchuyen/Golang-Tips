# Golang Functional Options Pattern

***Functional Options Pattern*** còn được gọi tắt là ***Options Pattern*** (không có func...)

`Functiona Options Pattern` là một dạng pattern (mẫu) để thiết kế 1 structure, giúp cho API có thể linh hoạt trong cấu hình và khởi tạo struct. Pattern này sử dụng kỹ thuật [Variadic function](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-function/Variadic-Functions.md) và [Interface](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/interface.md)
 
Trong [Ref.1], phần "Option 3: Functional Options Pattern", func New() được khai báo là 1 dạng `Variadic function`.

```
func New(options ...func(*Server)) *Server {
```

Ở đây:
- `options` là tham số đầu vào, nó là 1 dạng nhiều tham số (hiểu ở trình biên dịch Golang là) dấu `...`
- `func(*Server)` là 1 function có dạng/mẫu (pattern) đầu vào của func `New`.
- `func(*Server)` trả về chung 1 value là pointer `*Server`

### Tại sao lại cần Options Pattern
Trong [Ref.5] có đề cập tới vấn đề khi KHÔNG sử dụng `Options Pattern`
- Sẽ phải khai báo tất cả các value cho tham số đầu vào. Tưởng tượng nếu có tới > 10 tham số đầu vào cho 1 đối tượng (ví dụ: một thiết bị...) thì khai báo rất mệt mỏi, thậm chí bị nhầm lẫn.
- Hoặc nếu để giá trị là mặc định (khi khởi tạo đối tượng), thì sẽ phải làm thêm 1 method `setter` (set=tạo) để tạo giá trị cho 1 thuộc tính (tham số đầu vào).
- Khi 1 thuộc tính nào đó thay đổi (thêm, bớt, change-type,...) thì lại phải khai báo lại hàm khởi tạo (New-Object).
- Trật tự (argument ordering) của các tham số đầu vào không được đảm bảo (với trường hợp nhiều tham số), thậm chí còn phải chờ đợi nó được hoàn thiện rồi mới khởi tạo đối tượng.

### Không nhất thiết phải sử dụng Options Pattern
Trong [Ref.1] nói tới 3 cách để tạo ra 1 đối tượng, tất nhiên những đối tượng đơn giản thì nên dùng cách đơn giản, những đối tượng có cấu trúc xác định thì nên dùng `struct`.
> Một chú ý là khi sử dụng con trỏ `pointer` thì vẫn rẻ (inexpensive) (về memory). Vậy nên trong `Options Pattern` hay sử dụng con trỏ để khởi tạo và khai báo giá trị ban đầu.



### Những ưu điểm của Functional Options Pattern

Trong [Ref.2] có nói tới những ưu điểm của Options Pattern:
- Explicit: Tường minh. (MTC: thực sự thì đánh giá ưu điểm này có vẻ hơi khiên cưỡng, code clean thì cái gì cũng cần phải tường minh)
- Extensible: Mở rộng. (MTC: thực sự thì với mô tả trong tài liệu thì có vẻ không rõ ràng)
- Order of Arguments: Không cần quan trọng thứ tự của các tham số đầu vào. MTC: có vẻ đây là ưu điểm cụ thể nhất, trích dẫn là: `This gives us a lot of flexibility when compared to regular function arguments (which have to be in the correct order).` Nhưng với các tool IDE để lập trình ngày nay (Goland,...) thì có suggestion/auto-complete thì trật tự hay không vẫn giúp lập trình viên có thể hoàn thiện tốt các khai báo API.
- Without exposing fields/data to the outer world: tạm hiểu là không bị lộ dữ liệu ra bên ngoài. Trích dẫn trong 1 comment: `The main advantage of the functional options pattern is to allow you to set fields during initialization, ideally, without exposing those fields to the outer world.`
- Comment: `This pattern prevents the compiler from catching your misconfigurations.`


Vậy nên, có những ý kiến trái chiều về loại ***Options Pattern*** này  trong [Ref.4].

[Ref.1] https://golang.cafe/blog/golang-functional-options-pattern.html
[Ref.2] https://www.sohamkamani.com/golang/options-pattern/
[Ref.3] https://medium.com/@anar_py/option-pattern-in-go-944eda01677a
[Ref.4] https://www.reddit.com/r/golang/comments/pwjrjq/functional_option_pattern_is_terrible_in_practice/
[Ref.5] https://medium.com/@anar_py/option-pattern-in-go-944eda01677a
[Ref.6] https://levelup.gitconnected.com/options-pattern-in-golang-9a0384a9d8db

