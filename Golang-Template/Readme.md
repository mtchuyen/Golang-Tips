
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
[1] http://golangtutorials.blogspot.com/2011/06/go-templates.html

