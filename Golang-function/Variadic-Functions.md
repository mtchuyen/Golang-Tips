# Function tips for Golang

## Variadic Functions

***Dịch tạm***: Phân rã

### variadic(Adjective)

Taking a variable number of arguments; especially, taking arbitrarily many arguments.

(Dịch nôm na là: Lấy số lượng ***đối số (tham số đầu vào)*** thay đổi (số lượng tham số đầu vào không cố định); đặc biệt được dùng trong trường hợp lấy nhiều đối số tùy ý!)

***Origin***: variable + -adic

Ref: https://www.definitions.net/definition/variadic

### Variadic Functions

Ref: http://www.golangprograms.com/go-language/variadic-functions.html

***Variadic function*** gọi là 1 hàm mà ở đó chúng ta có thể truyền vào đối số (arguments) không giới hạn (infinite).

Để khai báo 1 ***variadic function*** thì:

- Biến (đối số) phân rã phải nằm cuối cùng của báo báo đầu vào hàm.
- Dùng dấu "..." (ellipsis) phía trước loại của biến (type of parameter)

Example: 

```
func calculation(str string, y ...int) int {
}
```
Tham khảo [link](https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085) giải thích rất chi tiết về cách thức hoạt động của hàm phân rã.
