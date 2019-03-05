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

#### Form action

```
<form action="/save/{{.Title}}" method="POST">
    <div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
    <div><input type="submit" value="Save"></div>
</form>
```

Thuộc tính `action="/save/{{.Title}}"` có nghĩa là khi bấm nút `submit` (với tên là `Save`) thì sẽ gọi đến handler *`save/<<title>>`*


### Template

Các web service thường phản hồi lại 1 request với dữ liệu (data) và html page, trong đó có nhiều nội dung được chuẩn hóa (ví dụ: header, footer, menu,...). Những dữ liệu sẽ được tùy chỉnh (modifications) dựa theo từng người dùng và loại request. Người ta gọi mẫu (Templates) là cách để gộp (merge) dữ liệu chung với dữ liệu riêng để đưa thành dữ liệu tùy chỉnh.

Trong Golang, thư viện ***template*** và các phương thức `Parse`, `ParseFile`, `Execute` để  lấy nội dung 1 template từ 1 chuỗi hoặc files, sau đó thực hiện merge các nội dung với nhau.

#### template packages

Có 2 thư viện quan trọng:
- https://golang.org/pkg/text/template/
- https://golang.org/pkg/html/template/

Các rules, schema sử dụng trong template (có thể được trình bày ở các phần dưới) đều được mô tả ở đây

#### Field substitution - {{.FieldName}}
Để đưa nội dung của 1 field vào trong template thì Golang định nghĩa 1 giá trị gọi là `trường thay thế (Field substitution)` được định nghĩa gồm:
- Kí hiệu bao hàm trường: Dấu ngoặc nhọn - curly bracket: `{{...}}`
- Tên trường: Viết hoa chữ cái đầu tiên với dấu chấm `.` (dot) phía trước (để biến này có thể exported).

Ví dụ: `"hello {{.Name}}!"`

```
package main

import (
    "os"
    "text/template"
)

type Person struct {
    Name string //exported field since it begins with a capital letter
}

func main() {
    t := template.New(“hello template”) //create a new template with some name
    t, _ = t.Parse("hello {{.Name}}!") //parse some content and generate a template, which is an internal representation

    p := Person{Name:"Mary"} //define an instance with required field
 
    t.Execute(os.Stdout, p) //merge template ‘t’ with content of ‘p’
}
```

### Referral
[1] https://gowebexamples.com/http-server/
[2] https://hackernoon.com/golang-template-1-bcb690165663
[3] https://viblo.asia
[4] http://www.cihanozhan.com/building-login-and-register-application-with-golang/
[5] https://mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout/
[6] http://golangtutorials.blogspot.com/2011/06/go-templates.html

