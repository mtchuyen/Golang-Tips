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

## 5 advanced testing techniques in Go

https://segment.com/blog/5-advanced-testing-techniques-in-go/

