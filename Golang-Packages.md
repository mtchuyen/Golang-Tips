# Packages tips for Golang

### Cấu trúc gói (packages)

```
package-name
── string
|  ├── case.go
|  ├── trim.go
|  └── misc.go
── number
   ├── arithmetics.go
   └── primes.go
```

Mỗi package nên được để bên trong 1 thư mục, ví dụ: package có tên là `number` được đặt trong thư mục `number`, tên package sẽ cùng tên với thư mục

### Package đặc biệt

- ***main*** là một package đặc biệt, mỗi chương trình (được build thành binary) phải có package này để  lệnh `go install` biết cách tạo ra executed file.
    + package `main` là entry point khi 1 chương trình có nhiều package được khai báo sử dụng
    + trong package `main` phải có function `main()` là entry point (nơi bắt đầu được chạy) của binary file.


### Go: What is the Unsafe Package?

[Go: What is the Unsafe Package?](https://medium.com/a-journey-with-go/go-what-is-the-unsafe-package-d2443da36350)


### Referrel:
[1] [Everything you need to know about Packages in Go](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)
