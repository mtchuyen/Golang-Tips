[Go: How Does a Goroutine Start and Exit?](https://medium.com/a-journey-with-go/go-how-does-a-goroutine-start-and-exit-2b3303890452)

[What Does a Goroutine Switch Actually Involve?](https://medium.com/a-journey-with-go/go-what-does-a-goroutine-switch-actually-involve-394c202dddb7)

[Understanding goroutines and channels](https://medium.com/@akankshadokania/why-golang-6b1f1c957dbd)


### How are Go channels implemented?

https://stackoverflow.com/questions/19621149/how-are-go-channels-implemented

The source file for channels is (from your go source code root) in /src/pkg/runtime/chan.go.

hchan is the central data structure for a channel, with send and receive linked lists (holding a pointer to their goroutine and the data element) and a closed flag. There's a Lock embedded structure that is defined in runtime2.go and that serves as a mutex (futex) or semaphore depending on the OS. The locking implementation is in lock_futex.go (Linux/Dragonfly/Some BSD) or lock_sema.go (Windows/OSX/Plan9/Some BSD), based on the build tags.

Channel operations are all implemented in this chan.go file, so you can see the makechan, send and receive operations, as well as the select construct, close, len and cap built-ins.

For a great in-depth explanation on the inner workings of channels, you have to read [Go channels on steroids](https://docs.google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub) by Dmitry Vyukov himself (Go core dev, goroutines, scheduler and channels among other things).

Here is a good talk that describes roughly how channels are implemented:
https://youtu.be/KBZlN0izeiY

Talk description:
> ***GopherCon 2017: Kavya Joshi - Understanding Channels***
> Channels provide a simple mechanism for goroutines to communicate, and a powerful construct to build sophisticated concurrency patterns. We will delve into the inner workings of channels and channel operations, including how they're supported by the runtime scheduler and memory management systems
> 
> ***Sub-title***: [Source code analysis of golang channel](https://laptrinhx.com/source-code-analysis-of-golang-channel-1392165230/)
> 

#### Channel in Go (Golang)

https://golangbyexample.com/channel-golang/


#### Diving Deep Into The Golang Channels.

https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8

#### Go: Buffered and Unbuffered Channels

https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268
