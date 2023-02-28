## Tables:

[1. First step](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#1-first-step)

[1.1. what-is-context](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#what-is-context-)

[1.2. why-context-is-important](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#why-context-is-important-)

[1.3. why-do-we-need-the-context-module](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#why-do-we-need-the-context-module)

[1.4. how-context-work](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#how-context-work-)

[1.5. how-to-create-a-context](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#how-to-create-a-context-)

[1.6. derived-contexts](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#derived-contexts)


[2. Context-is-immutable](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#2-context-is-immutable)

[2.1. method-withvalue](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#method-withvalue)

[2.2. access-parent-context-value](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#access-parent-context-value)


[3. Context-method](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#3-context-method)

[3.0.1. Context module được sử dụng trong ngữ cảnh nào?](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#01-context-module-%C4%91%C6%B0%E1%BB%A3c-s%E1%BB%AD-d%E1%BB%A5ng-trong-ng%E1%BB%AF-c%E1%BA%A3nh-n%C3%A0o)

[3.0.2. Cấu trúc Context interface](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#02-c%E1%BA%A5u-tr%C3%BAc-context-interface)

[3.1. Listening to a Cancellation Event](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#1-listening-to-a-cancellation-event)

[3.1.1 Listening to the Done() channel](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#11-listening-to-the-done-channel)

[3.1.2. Checking the Error From Err()](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#12-checking-the-error-from-err)

[3.2. Emitting a Cancellation Event](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#21-withcancel)

[3.2.1. Method: WithCancel](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#21-withcancel)

[3.3. Note](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#note)

## 1. First step

Đọc trong tài liệu này: 
[Ref1]: https://www.practical-go-lessons.com/chap-37-context

Bài viết giải thích rất chi tiết nguồn gốc, cách sử dụng (các ví dụ rất ngắn gọn). ***Tuy nhiên*** dài quá, chúng ta sẽ note tạm ở dưới đây để `bắt đầu với contexts`

Để ngắn hơn, chúng ta lấy từ bài viết này:
[Ref2]: https://blog.devgenius.io/go-introduction-to-context-package-2e5c844e1ae3

### What is Context ?
Context is a data that carries ***value, cancel signal, timeout signal, & Deadline signal***. With Context it would be easier for us to carry on value or signal between process.

### Why Context is Important ?
With context, when we want to cancel all of the process we just need to send a signal to context. So Automatically all process will be cancelled. Almost every part of GO-lang is using context such as database, http server, http client, etc. Even in Google, when develop using Go, Context is required and always send to every function.

#### Why Do We Need the Context Module?

[Ref4: Understanding Context in Golang](https://betterprogramming.pub/understanding-context-in-golang-7f574d9d94e0)

Imagine being the person taking orders in a restaurant.

When an order arrives, you *delegate* it to one of your many chefs.

What would you do if the customer *suddenly* decides to walk away?

Without a doubt, you will *stop* your chef from further processing the order to prevent any waste of ingredients!

That’s what the ***context*** module does!

It’s an argument passed to your functions and Goroutines and allows you to stop them promptly should you not require them anymore.


### How Context Work ?

For example if you want to cancel some process from Process A and send to another Process, All of another process will detect the signal which send from Process A and cancel the process.

![example of Context work](https://github.com/mtchuyen/Golang-Tips/blob/master/statics/Context-example-of-Context-work.png)


Và chúng ta sẽ quan tâm:
- how to create the contexts `(Background and TODO) `
- how to derive contexts `(WithValue, WithCancel, Deadline, and Timeout)`

### How to Create a Context ?
Because Context is an Interface, to create a context we need a struct which fit to Contract Interface. But in Context Package there is any function which we can use to create context instead of creating context manually.

There are 2 functions to create a context: (trong thư viện chuẩn của Golang, có 2 cách tạo ra Context)
1. ***context.Background()***

Create empty context, never cancelled, never timeout, and has no any value .

2. ***context.TODO()***

Create empty context like Background(), but usually used when we are not clear of the context that we want to use.

```
	//create empty context using Bakcground()
	background := context.Background()
	fmt.Println("backgr", background)

	// create empty context using TODO()
	todo := context.TODO()
	fmt.Println("todu:", todo)
```
This is the example to initiate ***empty*** context in go. 

### derived contexts
derived được hiểu là dẫn xuất - được tạo ra, sinh ra từ một cái khác.

Quá trình derived sẽ tạo ra một Context Tree (hay còn gọi là cây phả hệ Parent-Child-Subchild...)

https://btree.dev/golang-context

Việc dẫn xuất này giúp chúng ta:
- When a Context is canceled, all Contexts derived from it are also canceled.

parent context là 1 context gốc (root context)

derived contexts được hiểu là *childs* của ***Root-context***.
- Mỗi *child* sẽ được thêm các thông tin khác và được xử lý ở 1 layer khác.
---> đó chính là mối quan hệ ***Parent and Child Context***

#### how to derive contexts
Có vài phương thức hay sử dụng: WithValue, WithCancel, Deadline, and Timeout.

#### Parent and Child Context
Context is adopted ***Parent and Child*** concept. It’s means, when you are create a context you can create a child context from the existing context. Parent context is possible to have many child context, but child context is only allowed to have one parent context. This concept is similar to inheritance concept in Object Oriented Programming.

***Relation between Parent and Child Context***


![Relation between Parent and Child Context](https://github.com/mtchuyen/Golang-Tips/blob/master/statics/Context-Relation-between-Parent-and-Child-Context.png)


Parent & Child context will always be connected to each other. When you do something to Context A, all of Child and Sub child of context A will be impacted, but it wouldn’t impacted to Context B which in different inheritance path. You can access your data in parent Context from their own Child & Sub child context. but is not allowed to access data from another inheritance path context.

## 2. Context is Immutable

- ***Immutable*** hiểu nôm na là không thể thay đổi còn ***mutable*** là có thể thay đổi.
- 2 khái niệm Immutable và mutable thường được dùng class, object.
- Immutable Object: khi khởi tạo 1 đối tượng, thì trạng thái của tối tượng đó không thể thay đổi được sau khi việc khởi tạo đối tượng thành công. (Điều này có nghĩa là, bạn chỉ có thể get mà không thể set).

***Tại sao Context lại là Immutable***: sử dụng immutable sẽ tránh được sự thay đổi lẫn nhau khi đa tham chiếu.
- Sử dụng các immutable object làm tham số của method sẽ không sợ nó bị thay đổi sau khi method kết thúc
- Không có các method làm thay đổi trạng thái của các field (Ví dụ: chỉ có hàm get, không có các hàm set)
- Nếu có field nào là Object thì field đó cũng phải là 1 immutable Object hoặc khi khởi tạo/lấy ra field đó ta phải clone ra 1 bản khác.

### Method: WithValue
By default context is doesn’t have a value. You can add a value with pair (key-value) to the context. 

When you added new value to the context, a new child context will be created automatically. It’s means the original context doesn’t change. To add value to the context you can use function.

```
context.WithValue(parent, key, value)
```
Ví dụ:
```
contextA := context.Background()

contextB := context.WithValue(contextA, "b", "B")

fmt.Println("no.1", contextA)
fmt.Println("no.2", contextB)
fmt.Println("no.3", contextA)
```

- contextA là Parent, contextB là child.
- sau khi add value cho contextA (với method `WithValue`) thì contextA tự động clone ra 1 object khác là contextB = `new child context`.
- `original context ` = contextA. Print lần 1 với lần 3 thì value của contextA vẫn không thay đổi.


Try to run the above example code and see the result, you will see in your console that every child context will be connected to its own parent.

```
no.1 context.Background
no.2 context.Background.WithValue(type string, val B)
```
- `context.Background` là parent.
- contextB được nối với parent: `context.Background.WithValue(type string, val B)` = contextA`.``WithValue(type string, val B)`

### Access Parent Context Value
Child context can access the value of their own parent context

```
contextD := context.WithValue(contextB, "d", "D")
```
Ta có mối quan hệ: 

`[Parent]contextA --> [Child]contextB --> [Sub-Child]contextD`

Vậy contextD có thể access vào value của contextB

```
fmt.Println(contextD.Value("d"))
fmt.Println(contextD.Value("b"))
```

- The first line of code is try to access its own value, exactly you can access value from its own context.
- The second line of code is try to access its own parent context, this method is still allowed, you can access its own parent value.

## 3. Context method

Phần này sẽ tập trung vào nội dung của [Ref4: Understanding Context in Golang](https://betterprogramming.pub/understanding-context-in-golang-7f574d9d94e0)

Cơ bản:
- https://www.sohamkamani.com/golang/context-cancellation-and-values/
- https://www.prakharsrivastav.com/posts/golang-context-and-cancellation/


#### 0.1. Context module được sử dụng trong ngữ cảnh nào?

Xử lý các vấn đề when a client terminates the connection with a server.

|--> Điều gì sẽ xảy ra if the termination occurs while the server is in the middle of some heavy lifting work or database query?

|--> |--> The context module allows these processes to be stopped instantly as soon as they are not further in need.

Điều này cũng được mô tả qua ví dụ cụ thể của [Ref5: Go Context timeouts can be harmful](https://uptrace.dev/blog/golang-context-timeout.html)

The usage of the context module boils down to three primary parts:
1. Nhận sự kiện (thụ động): Listening to a cancellation event
2. Phát sự kiện (chủ động): Emitting a cancellation event
3. Passing request scope data: [Method: WithValue](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-basic/Context.md#method-withvalue)

#### 0.2. Cấu trúc Context interface

Trước khi đi vào các phần trên, ta xem cấu trúc của một `Context interface`
```
type Context interface {
    Done() <- chan struct{}
    Err() error
	
    Deadline() (deadline time.Time, ok bool)
    Value(key interface{}) interface{}
}
```
- The Done() function returns a channel that receives an `empty struct` when a context is cancelled.

- The Err() function returns a non-nil error in the `event of cancellation` otherwise, it returns a nil value.


### 1. Listening to a Cancellation Event

Có 2 cách để thực hiện Listening
- Cxt.Done()
- Cxt.Err()

#### 1.1 Listening to the Done() channel

```
func main() {
	http.ListenAndServe(":8000", http.HandlerFunc(handler))
}

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	select {
	case <-time.After(2 * time.Second):
		timnw := time.Now().String()
		w.Write([]byte("request processed: " + timnw))

	case <- ctx.Done():
		fmt.Println("request cancelled")
		return
	}
}
```
In the example above, we simulated a web service handler.

We used `time.After()` to simulate a function that takes two seconds to process a request.
- Khi client gửi request tới, server sẽ delay 2 second rồi mới trả về giá trị `request processed: 2021-02-27 15:04:14.0...` cho client
- Khi client gửi request tới và bấm hủy (chưa tới 2 second) thì server sẽ in ra `request cancelled` ở màn hình log của server.

Khi client hủy lệnh, `ctx.Done()` channel sẽ nhận được một `empty struct` để Go biết rằng context này đã hủy và tự động thực hiện đoạn lệnh bên trong (print... & return)

#### 1.2. Checking the Error From Err()

You can ***check for errors*** from `ctx.Err()` before executing some critical logic.

Ở ví dụ trên, thay vì ta thực hiện `fmt.Println("request cancelled")`, không có thông tin gì, ta sẽ làm cụ thể hơn (print ra nhiều thông tin hơn: lý do tại sao context bị cancel)

```
	case <- ctx.Done():
                fmt.Println(ctx.Err().Error())
                return
```


### 2. Emitting a Cancellation Event
Có 3 cách (method) hay được sử dụng (***WithXYZ***):

```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

Calling the ***cancelFunc*** emits an empty struct to the `ctx.Done()` channel and notifies downstream functions that are listening to it.

As you call the `WithXYZ` functions, they accept a parent context and return a new copy of the parent with a new `.Done()` channel.

```
rootCtx := context.Background()

child1Ctx, cancelFunc1 := context.WithCancel(rootCtx)
child2Ctx, cancelFunc2 := context.WithCancel(rootCtx)

child3Ctx, cancelFunc3 := context.WithCancel(child1Ctx)
```
When we call `cancelFunc1`, we will cancel `child1Ctx` and `child3Ctx`, while leaving `child2Ctx` unaffected.

***Different between WithTimeout and WithCancel***

- `WithTimeout`: query runs long time.
- `WithCancel`:  application is crashed/stopped/cancelled (no matter how long the query runs for).

***Different between WithTimeout and WithCancel***
- `WithTimeout`: `context.WithTimeout(ctx, time.Second*2)` --> 2 second wait
- `WithDeadline`: `context.WithDeadline(ctx, time.Now().Add(time.Second*20))`, current is `2021-02-28 16:02:06.206` and wait to `2021-02-28 16:02:26.206`

See: 
- http://www.inanzzz.com/index.php/post/olfs/cancelling-database-queries-with-context-withtimeout-and-withcancel-in-golang
- https://stackoverflow.com/questions/56721676/different-about-withcancel-and-withtimeout-in-golangs-context

#### 2.1. WithCancel

Source: 
- https://www.sohamkamani.com/golang/context-cancellation-and-values/
- https://viblo.asia/p/golang-context-cancel-va-cach-su-dung-gGJ59rVDKX2

```
func func1(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("=func1=: failed")
}

func func2(ctx context.Context) {
	fmt.Println("func2: exe operation2")
	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("func2: done")
	case <-ctx.Done():
		fmt.Println("func2: halted operation2")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := func1(ctx)
		fmt.Println("operation1 err:", err)
		if err != nil {
			cancel()
		}
	}()

	func2(ctx)
}
```
Case này, ở `func2` sẽ delay 50ms `case <-time.After(50 * time.Millisecond):` nên `func1` không được call

Khi tăng lên 500ms `case <-time.After(500 * time.Millisecond):` thì `func1` sẽ hoàn thành trước `func2`. Khi đó `cancel()` sẽ được gọi, nó sẽ truyền 1 signal (`empty struct`) để đoạn này `case <-ctx.Done():` sẽ được xử lý.

### Note
#### note 1: 
Always ***Defer Cancel Function***

```
    ctx := context.Background()

    ctx, cancel := context.WithTimeout(
        ctx, 
        3*time.Second,
    )
    defer cancel()
```

#### note 2:
A child ***can't and shouldn't*** cancel a context, it's the parent's call. 

See: https://stackoverflow.com/questions/66679577/cancel-context-from-child?rq=1

> you can use a channel to communicate, parent can listen to the channel, once there is an error, parent can cancel all children task.

Sometimes, calling `parentCancel()` doesn't mean ***everything will stop***. 

See: https://stackoverflow.com/questions/53009084/parent-child-context-cancelling-order-in-go

