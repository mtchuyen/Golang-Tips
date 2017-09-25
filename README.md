# Golang-Tips
Tips in Golang programming

# Goroutine

***1. Check number goroutine is running***

```go
type CountWG struct {
  sync.WaitGroup
  Count int // Race conditions, only for info logging.
}
// Increment on Add
func (cg CountWG) Add(delta int) {
  atomic.AddInt32(&cg.Count, delta)
  cg.WaitGroup.Add(delta)
}
// Decrement on Done
func (cg CountWG) Done() {
  atomic.AddInt32(&cg.Count, -1)
  cg.WaitGroup.Done()
}
```

# Golang Powered Robotics
https://gobot.io/: Gobot is a framework for robots, drones, and the Internet of Things (IoT), written in the Go programming language
