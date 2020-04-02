## package `sync`

 package sync giúp các hoạt động bất đồng bộ
 
 ***sync.Mutex***: Khóa để xử lý xong tác vụ
 ```go
 m := new(sync.Mutex)
 //some another code
 m.Lock()
 //some action --> another action must stop until thí action finish
 defer m.Unlock()
 ```
 
 ***sync.Atomic***: tương tác với số (integer, float)
 ```go
 atomic.AddInt32(&v, 1)
 ```
***sync.Once***: Khi bạn có một xử lý mà chỉ muốn thực hiện một lần duy nhất

***sync.Cond***: 

***sync.Pool***:

## Ref:
- [A Closer Look at Go’s sync Package](https://medium.com/@teivah/a-closer-look-at-go-sync-package-9f4e4a28c35a)
