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

# Golang Powered Robotics
https://gobot.io/: Gobot is a framework for robots, drones, and the Internet of Things (IoT), written in the Go programming language
