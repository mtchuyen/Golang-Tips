
### [Go: Finalizers](https://medium.com/a-journey-with-go/go-finalizers-786df8e17687)

Go runtime provides a method `runtime.SetFinalizer` that allows developers to attach a function to a variable that will be called when the garbage collector sees this variable as garbage ready to be collected since it became unreachable.
