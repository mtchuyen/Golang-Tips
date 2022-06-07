# Golang-Tips
Tips in Golang programming
### Go design

https://go.googlesource.com/proposal/+/master/design

[The Top 10 Most Common Mistakes I’ve Seen in Go Projects](https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65)

### Go Architecture
Những nguyên lý thiết kế app trên ngôn ngữ Golang:

#### Hexagonal Architecture in Go

Đọc cơ bản về kiến trúc ***lục giác (Hexagonal)***, còn gọi là kiến trúc ***Ports and Adapter*** trong [link](https://github.com/mtchuyen/back-end/blob/master/Hexagonal-Architecture.md)


[Hexagonal Architecture in Go](https://medium.com/hexagonal-architecture-in-go/hexagonal-architecture-in-go-94f4ed15392a)

[Hexagonal Architecture in Go](https://blog.devops.dev/hexagonal-architecture-in-go-58dd2386dea7)


#### [Clean Architecture with Go](https://medium.com/@martinezdelariva/clean-architecture-with-go-60feb7aac3f8)

#### Nên tuân thủ 1 số Pattern:
- [Patterns in Go applications](https://medium.com/@tranngoclam/patterns-in-go-applications-42dcd10fd0e5)
- [Practical SOLID in Golang: Interface Segregation Principle](https://levelup.gitconnected.com/practical-solid-in-golang-interface-segregation-principle-f272c2a9a270)


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

## Auto-Scaling and Self-Defensive Services in Golang

https://dzone.com/articles/auto-scaling-and-self-defensive-services-in-golang-1

### Server-Side Tracking Without Cookies In Go
https://medium.com/dev-genius/server-side-tracking-without-cookies-in-go-44f3703331ba

## Understanding Go Lang Memory Usage

https://deferpanic.com/blog/understanding-golang-memory-usage/

## Integration of a Go service with systemd
[1](https://vincent.bernat.im/en/blog/2017-systemd-golang), [2](https://vincent.bernat.im/en/blog/2018-systemd-golang-socket-activation)

# Internet of Things (IoT)
## Simple IoT Messages Delivery With GoLang

https://ralvescosta.medium.com/simple-iot-messages-delivery-with-golang-1-b1ea64d7d3ae

https://ralvescosta.medium.com/simple-iot-messages-delivery-with-golang-2-9d9ebfb75fd4


## Golang Powered Robotics
https://gobot.io/: Gobot is a framework for robots, drones, and the Internet of Things (IoT), written in the Go programming language

# Tiền điện tử

https://vicrypto.org/

# Best Golang Books

### [Awesome Go Books](https://github.com/dariubs/GoBooks)


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

# awesomeness

[Companies Using Golang by Domain — Golang Use Cases](https://medium.com/@softkraft/companies-using-golang-by-domain-golang-use-cases-6870b9001e82?)

## Generated awesomeness

[Collection of useful resources on Bioinformatics, data science, machine learning, programming language (Python, Golang, R, Perl, etc.) and miscellaneous stuff.] (https://github.com/shenwei356/awesome)

## Awesome Go

https://github.com/avelino/awesome-go

## Golang Libraries

https://github.com/dreadl0ck/golang-libs

## Awesome Go performance

https://github.com/cristaloleg/awesome-go-perf

## Awesome Go Tools

https://github.com/gobuild/awesome-go-tools

## Golang Web Scraping

https://github.com/lorien/awesome-web-scraping

[How Golang is Used for Web Scraping with Concurrency?](https://iweb-scraping-services.medium.com/how-golang-is-used-for-web-scraping-with-concurrency-e9873cfd935a)


## algorithms with golang solution

https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-Algorithms/Readme.md


## awesome libraries for golang, packages and resources for Quants (Quantitative Finance)

https://github.com/goex-top/awesome-go-quant

# Golang Courses
[9 Best Golang Programming Courses for Beginners to Learn in 2022](https://medium.com/javarevisited/what-is-go-or-golang-programming-language-why-learn-go-in-2020-1cbf0afc71db)

## Dạy Golang bằng tiếng Việt
- Lập trình Go: https://phocode.com/blog/2016/08/26/go-con-tro/
- https://techblog.vn/tags/golang


# Referral:
- [Cách sử dụng package `sync` của golang](https://kipalog.com/posts/Cach-su-dung-package--sync--cua-golang)


# Secure Coding
- Basic: https://owasp.org/www-pdf-archive/OWASP_SCP_Quick_Reference_Guide_v2.pdf
- [Introduction to secure coding in Golang](https://mykparmar007.medium.com/introduction-to-secure-coding-in-golang-f229c6668c25)
- [5 Secure Coding Tips in Go](https://medium.com/picus-security-engineering/5-secure-coding-tips-in-go-a3e5ec23d7fd)
- [Hacking with Go](https://github.com/parsiya/Hacking-with-Go)
- [Top 6 security best practices for Go](https://blog.sqreen.com/top-6-security-best-practices-for-go/)
- [TLS validation: implement OCSP and CRL verifiers in Go](https://www.cossacklabs.com/blog/tls-validation-implementing-ocsp-and-crl-in-go/)


## vulnerabilities
- [Scanning Go dependencies for vulnerabilities](https://jcdan3.medium.com/scanning-go-dependencies-for-vulnerabilities-b82db3d56b27)
- [Scanning Go source for vulnerabilities with gosec](https://systemweakness.com/scanning-go-source-for-vulnerabilities-5f29773ecc9d)



 
