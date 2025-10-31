## Basic
Should be read first: (Goroutine)[https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Goroutine.md] 

### What are Channels?
Channels are the Go way to communicate between ***goroutines***. They let goroutines send and receive data safely.

```go
package main

import "fmt"

func worker(doneGR chan bool) { // doneGR is a channel that accepts boolean
	fmt.Println("Worker: Doing some work...")
	//time.sleep(1)
	// Simulate work
	fmt.Println("Worker: Done!")
	doneGR <- true // Send a true signal to the 'doneGR' channel
}
//doneGR: done GoRoutine
//doneM: done in Main
func main() {
	// Create a channel that sends booleans
	doneM := make(chan bool)

	// Start the worker goroutine, passing the channel
	go worker(doneM)

	// Block until we receive a value from the 'done' channel
	<-doneM
	fmt.Println("Main: Worker finished!")
}

```

Channels are typed (e.g., `chan bool`, `chan string`).

### What’s buffered vs unbuffered channels?
- ***Unbuffered Channel*** (default): `make(chan int)`. Sender waits until a receiver is ready, and receiver ***waits*** until a sender is ready. It's like a *handshake*.
- ***Buffered Channel***: `make(chan int, 5)`. Can hold up to 5 values ***before*** the sender blocks. It's like a queue.

```go
func main() {
 // Buffered channel with capacity 2
 messages := make(chan string, 2)
 messages <- "hello"     // Sends immediately, doesn't block
 messages <- "world"     // Sends immediately, doesn't block
 
 fmt.Println(<-messages) // Receive "hello"
 fmt.Println(<-messages) // Receive "world"
}
```

Use buffered channels when you want to decouple sender and receiver a bit, allowing a few messages to queue up.

## Referral:

[1]: [Go Simplified: Channels](https://medium.com/@deepakschoudhary/go-simplified-channels-c480f5771343)

