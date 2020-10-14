### Handler
#### Xử lý URL Path
`viewHandler` cho phép người dùng xem một trang của page. Nó sẽ xử lý URL với tiền tố `/view/`.

Ví dụ: `example.com/view/contentexample` thì sẽ xử lý path: `/view/contentexample`

```
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    fmt.Fprintf(w, "<h1>Test</h1><div>%s</div>", title)
}
```
Đường dẫn được cắt lát lại với `[len("/view/"):]` để bỏ `/view/` thành component của request path.

#### Xử lý lỗi
```
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
```

Hàm `http.Error` gửi một mã phản hồi HTTP được chỉ định (trong trường hợp này là "Internal Server Error") và thông báo lỗi.

#### Redirect sau khi đã xử lý nghiệp vụ

```
    func saveHandler(w http.ResponseWriter, r *http.Request) {
        //some action

        http.Redirect(w, r, "/view/"+title, http.StatusFound)
    }
```

Trong nghiệp vụ `saveHandler`, sau khi xử lý dữ liệu và thực hiện việc lưu dữ liệu, cần chuyển hướng về page có nội dung vừa `save`, bằng cách thực hiện chỉ hướng `http.Redirect`.

#### Assets and Files (Static files)

```
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
```
Ví dụ: Nếu có query: `example.com/static/soha.gif` vào thư mục `assets/` lấy file `soha.gif`.

#### Form action

```
<form action="/save/{{.Title}}" method="POST">
    <div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
    <div><input type="submit" value="Save"></div>
</form>
```

Thuộc tính `action="/save/{{.Title}}"` có nghĩa là khi bấm nút `submit` (với tên là `Save`) thì sẽ gọi đến handler *`save/<<title>>`*

### Benchmark

#### hey is a tiny program that sends some load to a web application.

Link: https://github.com/rakyll/hey/


### Serving up HTML with Go

https://blog.meshstudio.io/serving-up-html-with-go-92f767856daf


### Automagical HTTPS with Docker and Go

***Guide https free by Go***

- https://github.com/kjk/go-cookbook/tree/master/free-ssl-certificates
- https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go.html

And simpler:

- https://medium.com/weareservian/automagical-https-with-docker-and-go-4953fdaf83d2



### Web Service Architecture for Golang Developers

https://medium.com/@boobo94/web-service-architecture-for-golang-developers-a4147b8141ff

### Learning Golang through WebAssembly

[Learning Golang through WebAssembly](https://www.aaron-powell.com/posts/2019-02-05-golang-wasm-2-writing-go/)

https://dev.to/hajimehoshi/gopherjs-vs-webassembly-for-go-148m

[Some notes about the upcoming WebAssembly support in Go](https://blog.owulveryck.info/2018/06/08/some-notes-about-the-upcoming-webassembly-support-in-go.html)

https://matthewphillips.info/programming/wasm-golang-ce.html

https://www.sitepen.com/blog/compiling-go-to-webassembly/

https://mycodesmells.com/post/building-wasm-applications-in-go-1-11

### Streaming Images from MongoDB using Golang
https://medium.com/adam-on-software-engineering/streaming-images-from-mongodb-using-golang-651fc2b8384e

### Referral
[1] https://gowebexamples.com/http-server/

[2] https://hackernoon.com/golang-template-1-bcb690165663

[3] https://viblo.asia

[4] http://www.cihanozhan.com/building-login-and-register-application-with-golang/

[5] https://mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout/
