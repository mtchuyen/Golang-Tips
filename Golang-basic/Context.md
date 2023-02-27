## First step

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

## Context is Immutable

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
