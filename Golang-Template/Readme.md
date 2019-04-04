
## Template

Các web service thường phản hồi lại 1 request với dữ liệu (data) và html page, trong đó có nhiều nội dung được chuẩn hóa (ví dụ: header, footer, menu,...). Những dữ liệu sẽ được tùy chỉnh (modifications) dựa theo từng người dùng và loại request. Người ta gọi mẫu (Templates) là cách để gộp (merge) dữ liệu chung với dữ liệu riêng để đưa thành dữ liệu tùy chỉnh.

Trong Golang, thư viện ***template*** và các phương thức `Parse`, `ParseFile`, `Execute` để  lấy nội dung 1 template từ 1 chuỗi hoặc files, sau đó thực hiện merge các nội dung với nhau.

Đọc [2] (hoặc [link](http://golang-examples.tumblr.com/post/87553422434/template-and-associated-templates)) để hiểu về cách thức hoạt động của template.

### template packages

Có 2 thư viện quan trọng:
- https://golang.org/pkg/text/template/
- https://golang.org/pkg/html/template/

Các rules, schema sử dụng trong template (có thể được trình bày ở các phần dưới) đều được mô tả ở đây

### Template Names

Tất cả các template đều phải có ***name***

- Template name được định nghĩa bằng `template.New()`
```
 t := template.New("main1") //name of the template is main1
```
- Templates có thể được tìm kiếm bởi ***Name*** khi sử dụng func `template.Lookup()`.
```
t = t.Lookup("ROOT") // <-- here
if t == nil {
    return fmt.Errorf("ROOT template not found in %v", set)
}
```
- Template name can be queried(`template.Name()`)
- 

### Parsing Templates

- ***template.Parse()***: để parse string
```
t, _ = t.Parse(tmpl) // parsing of template string: tmpl is a string
```

- ***template.ParseFiles()***: Giúp parser nhiều file, với các filename được chỉ định (tham số truyền vào) (nó là 1 `variadic function`)

- ***template.ParseGlob()***:  sử dụng `pattern matching` giúp parser tất cả các file có cùng định dạng.

Cả 3 phương thức trên đều có thể dùng hàm `Execute` để render 1 template, ngoài ra phương thức ***ParseFiles*** còn có thể dùng hàm `ExecuteTemplate` để render.

2 phương thức `PraseFiles` và `ParseGlob` luôn trả về `tempale name` đầu tiên, nên nếu muốn dùng `template name` về sau thì cần dùng phương thức `ExecuteTemplate`. Xem ví dụ `helloworld.go` để hiểu chi tiết hơn.

> The returned `template name` will be the first file in PraseFiles() and the first file matched in the regular expression in ParseGlob().

```
t, _ := template.ParseFiles("t1.html", "t2.html")
fmt.Println(t.Name())
```

Return:

> t1.html


### Executing Templates

- ***template.Execute()***:

```
t.Execute(w, "Asit") // --> luôn lấy template name đầu tiên của t để parser và print ra biến w
```

- ***template.ExecuteTemplate()***:

```
t.ExecuteTemplate(w, "t2.html", "Golang") // --> Lấy template name có tên là t2.html của t để parser và print
```

`template.ExecuteTemplate()` calls `template.Execute()` internally. It basically looks for the named template and executes that one.

### Field substitution - {{.FieldName}}
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
#### IF-ELSE
#### PRINT
#### DEFINE: a template Set
```
{{define "content1"}}
<p>content 1</p>
{{end}}
```
Key points to note:
* each template is enclosed within {{define}} and {{end}}
* each template is given a unique name - repeating the same name will cause a panic. Ex: "content1"
* Dùng 1 template khác bên trong 1 template hiện hành - using `{{template "template name"}}`

### Referral

[1] http://golangtutorials.blogspot.com/2011/06/go-templates.html

[2] Một bài viết giải thích hoạt động của template: http://golang-examples.tumblr.com/post/87553422434/template-and-associated-templates

[3] series golang template: ttps://www.calhoun.io/intro-to-templates-p1-contextual-encoding/

[4] series: [1](https://hackernoon.com/golang-template-1-bcb690165663), [2](https://hackernoon.com/golang-template-2-template-composition-and-how-to-organize-template-files-4cb40bcdf8f6)


