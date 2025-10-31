# Function tips for Golang

## How do define a function?
You use the func keyword. You specify input parameters and what it returns.

- A simple function that takes two ints and returns their sum
```
func add(a int, b int) int {
 return a + b
}
```
- A function that returns multiple values (very common in Go!)
```
func greet(name string) (string, int) {
 message := "Hello, " + name
 length := len(message)
 return message, length // Return two values
}
```
==> Then call in main
```
func main() {
 sum := add(5, 3)
 fmt.Println("Sum:", sum)
 greeting, len := greet("Charlie")
 fmt.Println(greeting, " (length:", len, ")")
}
```

## finalizers

### [Go: Finalizers](https://medium.com/a-journey-with-go/go-finalizers-786df8e17687)

Go runtime provides a method `runtime.SetFinalizer` that allows developers to attach a function to a variable that will be called when the garbage collector sees this variable as garbage ready to be collected since it became unreachable.


### finalizers in Go
- [Go: Finalizers](https://medium.com/a-journey-with-go/go-finalizers-786df8e17687)
- [An example of the use of a finalizer](https://gist.github.com/deltamobile/6511901)
- [Mystery of finalizers in Go](https://lk4d4.darth.io/posts/finalizers/)


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
