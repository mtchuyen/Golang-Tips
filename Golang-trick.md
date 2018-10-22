# Golang tricks

## Separate your binary from your application

Tham khảo: [Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)

Nếu để file `main.go` ở trong root path của project thì có đặc điểm

***Tiện dụng:***
- Ứng dụng để quản lý code, các thư viện được nhúng dễ theo cấu trúc logic
- Biên dịch đơn giản
- Khi ai đó lấy từ repo về bằng `go get` thì chương trình sẽ tự compile và cài đặt

***Khó dùng:***
- Ứng dụng khó được dùng như 1 thư viện (library)
- only have one application binary.

to fix this is to simply use a ***cmd*** directory in my project where each of its subdirectories is an application binary.

```go
adder/
  adder.go
  cmd/
    adder-client/
      main.go
    adder-server/
      main.go
```
Here we have 2 separate application binaries that can be built and is installed: adder-client, adder-server.

- `adder.go` sẽ là thư viện được chia sẻ chung (ví dụ: package adder).
- `main.go` trong thư mục `adder-client` sẽ biên dịch thành file binary `adder-client`
- `main.go` trong thư mục `adder-server` sẽ biên dịch thành file binary `adder-server`

Như vậy việc chuyển file `main.go` khỏi thư mục root của project sẽ giúp project vừa có thể thành thư viện sử dụng chung, vừa có thể thành các ứng dụng khác nhau sử dụng thư viện  `adder` ví dụ như:
- adder-client: sẽ giúp người dùng tương tác bằng command
- adder-server: sẽ giúp người dùng tương tác bằng web.

## Don’t use global variables
Tham khảo: [Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)

Lý do: xem trong phần `tham khảo`

```go
type MyDB struct {
  *sql.DB
}

func(db *MyDB) handlerOne(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Do something databasey
  })
}

func(db *MyDB) handlerTwo() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Do something else databasey
  })
}

func main() {
  db, err := sql.Open("postgres", "…")
  // Handle
  mydb := &MyDB{db}
  http.Handle("/", myDB.handlerOne())
  // Etc.
}
```
Với kiểu khai báo biến MyDB là `struct` giúp việc truyền biến dễ dàng hơn.

## Variadic Functions

***Dịch tạm***: Phân rã

### variadic(Adjective)

Taking a variable number of arguments; especially, taking arbitrarily many arguments.

(Dịch nôm na là: Lấy số lượng ***đối số (tham số đầu vào)*** thay đổi (số lượng tham số đầu vào không cố định); đặc biệt được dùng trong trường hợp lấy nhiều đối số tùy ý!)

***Origin***: variable + -adic

Ref: https://www.definitions.net/definition/variadic

### Variadic Functions

Ref:
- [Variadic Functions in Golang](http://www.golangprograms.com/go-language/variadic-functions.html)

- [Variadic Functions in Programming]https://en.wikipedia.org/wiki/Variadic_function


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

[Bài viết](https://medium.com/golangspec/variadic-functions-in-go-13c33182b851) giới thiệu 1 ngữ cảnh sử dụng hàm phân rã bị sai, và phải sửa thành:
```
numbers := []int{1, 2, 3}
f(numbers...)
```
Như vậy dùng dấu 3 chấm `"..."` trong biến ``numbers...``` (dạng slice) sẽ phân rã biến ```numbers``` thành các ```number```


## 5 advanced testing techniques in Go

https://segment.com/blog/5-advanced-testing-techniques-in-go/

## Optimising Go allocations using pprof

https://www.robustperception.io/optimising-go-allocations-using-pprof/

## A Makefile for your Go project

https://vincent.bernat.im/en/blog/2017-makefile-build-golang


