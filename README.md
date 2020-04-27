# Golang-Tips
Tips in Golang programming
### Go design

https://go.googlesource.com/proposal/+/master/design

[The Top 10 Most Common Mistakes I’ve Seen in Go Projects](https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65)


# Golang Concepts

https://medium.com/jexia/master-the-world-of-golang-issue-1-4c4c3732e13f

## A Class Factory in Golang (Google Go)

http://www.minaandrawos.com/2014/09/25/class-factory-in-golang-google-go/

https://matthewbrown.io/2016/01/23/factory-pattern-in-golang/

Chúng ta xem xét 1 interface có tên là `AbstractFactory`:

```go
type AbstractFactory interface {
        CreateMyLove()
}
```
Ở đây có 2 khái niệm:
- Interface `AbstractFactory` có method là `CreateMyLove()`
- `CreateMyLove()` được gọi là 1 ***factory method***.

## finalizers in Go
- [Go: Finalizers](https://medium.com/a-journey-with-go/go-finalizers-786df8e17687)
- [An example of the use of a finalizer](https://gist.github.com/deltamobile/6511901)
- [Mystery of finalizers in Go](https://lk4d4.darth.io/posts/finalizers/)

## Auto-Scaling and Self-Defensive Services in Golang

https://dzone.com/articles/auto-scaling-and-self-defensive-services-in-golang-1

## Understanding Go Lang Memory Usage

https://deferpanic.com/blog/understanding-golang-memory-usage/

## Integration of a Go service with systemd
[1](https://vincent.bernat.im/en/blog/2017-systemd-golang), [2](https://vincent.bernat.im/en/blog/2018-systemd-golang-socket-activation)


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
