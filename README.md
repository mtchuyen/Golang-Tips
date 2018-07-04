# Golang-Tips
Tips in Golang programming

### Golang embedded type
***Embedded Type*** là khai báo một type nằm trong một type khác nhưng ***không khai báo tên***, trường mà không khai báo tên còn được gọi là ***embedded field***. 

Ví dụ như sau: 
```go
type Author struct {
    AuthorName string
    AuthorAge int
}

type Post struct {
    Title string
    Content string
    Author        // ---> Embedded field
}
```
Trong ví dụ trên nếu chúng ta đặt tên cho trường *Author* như bình thường thì chúng ta sẽ có Struct lồng nhau, còn nếu dùng ***Embedded field*** thì chúng ta có thể coi như Struct *Post* có đầy đủ các trường của cả 2 Struct (tên trường không được phép trùng nhau)

```go
type Post struct {
    Title string
    Content string
    AuthorName string
    AuthorAge int
}
```
Bằng cách này chúng ta sẽ có thể sử dụng cả 2 Struct *Post* và *Author* mà không cần khai báo lại các trường trùng lặp.

Khi lấy dữ liệu cũng có thể lấy trực tiếp mà không cần qua Struct trung gian, ví dụ lấy tên tác giả thay vì ***post.Author.AuthorName*** thì ta chỉ cần ***post.AuthorName***.

```go
type Author struct {
	AuthorName string
	AuthorAge  int
}

type Post struct {
	Title   string
	Content string
	Author  // Embedded field
}

func main() {
	author := Author{
		AuthorName: "Robin",
		AuthorAge:  50,
	}
	post := Post{
		Title:   "Simple Post",
		Content: "This post has no content",
		Author:  author,
	}

	fmt.Println(post)
	fmt.Println(post.Title)
	fmt.Println(post.AuthorName)   // Don't need post.Author.AuthorName
}
```
Chi tiết bài viết ở đây: [Golang embedded type - Kế thừa trong Go](https://techmaster.vn/posts/34682/golang-embedded-type-ke-thua-trong-go)

### package `sync`
 package sync giúp các hoạt động bất đồng bộ
 
 ***sync.Mutex***: Khóa để xử lý xong tác vụ
 ```go
 m := new(sync.Mutex)
 //some another code
 m.Lock()
 //some action --> another action must stop until thí action finish
 defer m.Unlock()
 ```
 
 ***sync.Atomic***: tương tác với số (integer, float)
 ```go
 atomic.AddInt32(&v, 1)
 ```
***sync.Once***: Khi bạn có một xử lý mà chỉ muốn thực hiện một lần duy nhất

***sync.Cond***: 

***sync.Pool***:


# Golang Concepts

https://medium.com/jexia/master-the-world-of-golang-issue-1-4c4c3732e13f

## A Class Factory in Golang (Google Go)

http://www.minaandrawos.com/2014/09/25/class-factory-in-golang-google-go/

## Auto-Scaling and Self-Defensive Services in Golang
https://dzone.com/articles/auto-scaling-and-self-defensive-services-in-golang-1


# Golang Powered Robotics
https://gobot.io/: Gobot is a framework for robots, drones, and the Internet of Things (IoT), written in the Go programming language

# Tiền điện tử

https://vicrypto.org/

# Best Golang Books
- The Go Programming Language (advanced programmers)
> This is maybe one of the most complete books about go language. It covers the basics and then goes deeper in lower levels. The downside of this book is that it don’t cover the higher level of the language, like web.
If you want to master this language it’s a good choice to buy.

- Go in Practice: Includes 70 Techniques
> It approaches a more higher level of this language but also gives good foundations to be a go programmer. It covers web applications development, micro services and even deployment.
A good solution for the developer that want to use go for webapplications.

- Go Web Programming
> This book is more oriented and specific to web development. If you intend to build apis, develop full stack applications and understand how go fits exactly in this world, this book is an excellent choice.
If you want to master web development this is the book.

- Concurrency in Go: Tools and Techniques for Developers
> Concurrency is one of the most appealing features in go. Although the simplicity of concurrency primitives in go, building more complex things requires some knowledge and practice.
If you want to master concurrency go for this one.

- Go Programming Blueprints — Second Edition
> For me this is the go programming bible. If you read this book, you will have the necessary tool to build massive applications with go. Besides that, Mat Ryer is one of the most enthusiastic voices in go community.

# Dạy Golang bằng tiếng Việt
- Lập trình Go: https://phocode.com/blog/2016/08/26/go-con-tro/
- https://techblog.vn/tags/golang

# Referral:
- [Cách sử dụng package `sync` của golang](https://kipalog.com/posts/Cach-su-dung-package--sync--cua-golang)
