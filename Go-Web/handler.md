## Handler

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
