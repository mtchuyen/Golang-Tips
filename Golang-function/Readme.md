## How do define a function?
You use the func keyword. You specify input parameters and what it returns.

- A simple function that takes two ints and returns their sum
```
func add(a int, b int) int {
 return a + b
}
```
- A function that returns multiple values (very common in Go!)
```
func greet(name string) (string, int) {
 message := "Hello, " + name
 length := len(message)
 return message, length // Return two values
}
```
==> Then call in main
```
func main() {
 sum := add(5, 3)
 fmt.Println("Sum:", sum)
 greeting, len := greet("Charlie")
 fmt.Println(greeting, " (length:", len, ")")
}
```

## finalizers

### [Go: Finalizers](https://medium.com/a-journey-with-go/go-finalizers-786df8e17687)

Go runtime provides a method `runtime.SetFinalizer` that allows developers to attach a function to a variable that will be called when the garbage collector sees this variable as garbage ready to be collected since it became unreachable.


### finalizers in Go
- [Go: Finalizers](https://medium.com/a-journey-with-go/go-finalizers-786df8e17687)
- [An example of the use of a finalizer](https://gist.github.com/deltamobile/6511901)
- [Mystery of finalizers in Go](https://lk4d4.darth.io/posts/finalizers/)
