## regular expressions

https://zetcode.com/golang/regex/

https://yourbasic.org/golang/regexp-cheat-sheet/

https://www.golangprograms.com/find-dns-records-programmatically.html

https://geekflare.com/regular-expression-tester/

https://code.tutsplus.com/tutorials/8-regular-expressions-you-should-know--net-6149

https://code.tutsplus.com/tutorials/regular-expressions-with-go-part-1--cms-30403

https://stackoverflow.com/questions/9363145/regex-for-extracting-filename-from-path

https://stackoverflow.com/questions/47131996/go-how-to-check-if-a-string-contains-multiple-substrings/47133055

https://howkteam.vn/course/Gioi-thieu-cau-truc-vong-lap-trong-C/Regular-Expression-trong-C-1427

https://viblo.asia/p/regular-expressions-regex-khong-he-kho-nhu-nhung-gi-ban-thay-i-GrLZDAXOlk0

https://gpcoder.com/2222-huong-dan-su-dung-bieu-thuc-chinh-quy-regular-expression-trong-java/

https://topdev.vn/blog/regex-la-gi/

https://freetuts.net/cac-quy-tac-regular-expression-can-ban-tiep-theo-66.html

https://quantrimang.com/regex-trong-python-165471

#### Text boundary anchors

| Symbol |	Matches | 
| --- | --- |
| \A	| at beginning of text
| ^		| at beginning of text or line (`^m` bắt đầu của text phải là chữ `m`)
| $		| at end of text/string (`js$` kết thúc chuỗi phải là chữ `js`)
| \z | at end of text
| \b	| 	at ASCII word boundary
| \B	| 	not at ASCII word boundary

### Example: 
- `^[a-z]+$`: Kiểm tra một chuỗi là các ký tự thường.
- `^[0-9]+$`: Kiểm tra một chuỗi là số (bắt đầu `^` là số, chỉ có số, kết thúc chuỗi `$`cũng là số)
- 

#### Ex1: Check Filename use this regular expression

```
"^[a-zA-Z0-9_\\-\\.]*$"
```

Explanations:
```
^ : start of string
[ : beginning of character group
a-z : any lowercase letter
A-Z : any uppercase letter
0-9 : any digit
_ : underscore 
- : hyphen (double escape signal \)
. : dot  (double escape signal \)
] : end of character group
* : zero or more of the given characters
$ : end of string
```

***Example:***
```
 Check validate spoxxxx_adxgg1.js 
 Check validate iplay-mutex-sxx.js 
 Check validate 131.js 
 
 Check validate C:\AppServ\www\project\storage\pr.vn.js 
 Filename C:\AppServ\www\project\storage\pr.vn.js is invalid --> (\ signal in path is not in matching regular)
```
#### Replace match char

***Ex:*** `1%27-rt(doc.co)-%27`

Cần loại bỏ các dấu (ký tự) đặc biệt: `%`,...
```
reg, err := regexp.Compile("[^a-zA-Z0-9_\\-\\.]")
cleanString := reg.ReplaceAllString(orginString, "0")
```

Giải thích:
- `[...]`: group các ký tự cần tìm kiếm (để thay thế)
- `[^...]`: có dấu `^` để chỉ rằng các group ký tự bên trong là hợp lệ (chấp nhận). Gồm các ký tự từ 0 đến 9, a đến z, A đến Z, `_` (underscore), `-` (hyphen), `.` (dot)
- không có dấu `$` ở cuối như ở phần ***Check***

***Kết quả:*** 1`0`27-rt`0`doc.co`0`-027
- Dấu `%` bị thay thế bởi (số) `0`.
- Dấu `(`, `)` bị thay thế bởi (số) `0`.

| Origin | Result |
| --- | ---
| `1&27-alert*doc.co#-%27` | 1`0`27-alert`0`doc.co`0`-027
| `1%27-alert!doc@.co#-%27` | 1`0`27-alert`0`doc`0`.co`0`-`0`27
