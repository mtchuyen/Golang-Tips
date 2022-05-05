## interface

***Note***: trước khi đọc `Interface` cần nằm được `struct` (structure) và `method` (function của struct)

### Type: đại diện cho kiểu giá trị mà chúng ta sử dụng. 

***Tác dụng:***

- Nhờ có Type mà các compiler có thể xác định được một số lỗi trong quá trình compile.

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
