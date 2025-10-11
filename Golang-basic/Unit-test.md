## What are Unit Tests?
> In computer programming, ***unit testing*** is a software testing method by which individual units of source code — sets of one or more computer program modules together with associated control data, usage procedures, and operating procedures — are tested to determine whether they are fit for use

The ***unit test*** is the first level of software test and followed by integration tests and ui tests.

### Tác dụng Unit Test
- Unit testing aids in the discovery and elimination of issues early in the development cycle, as well as the prevention of regressions during refactoring. 
- A decent unit test may also act as documentation for new developers on the project.

![Test-Pyramid-UnitTest](https://github.com/mtchuyen/Golang-Tips/blob/44c79ef3287c8b44a8aff00c0266bcdaaeedadd5/statics/Test-Pyramid1_9sF9aGdGXWPm28Q8.png)

### Một số loại test được support trong Golang
- Test Suites: is a group of test cases on a particular feature or component. Use `Test Suites` to organize your test cases.
- Table Tests:  allow us to cover all the edge cases we expect to encounter. Use `Table Tests` to cover all your program's edge cases and avoid run-time crashes.

## Syntax

### Syntax for File

Test case are placed inside ***_test.go*** files

Ex:
- math_test.go
- ulti_comm_test.go


### Syntax for Func

`func TestXyy (t *testing.T)`

- `Test` is prefix
- `X` is upper-case char in ***Xyy*** func name

Ex: 
```
func TestGenGUID()(t *testing.T) {

}
```
### Syntax for Command

go test -run ***Func-name:Xyy***

Ex: `go test -run GenGUID`


***Run all Test func***
- `go test -v`


## Lần đầu tạo Unit Test in Go project

[Ref.4]

- ***Step 1: project structure***: Tạo ra 1 file code có hậu tố  ***_test.go***
- ***Step 2: make simple func***: Tạo ra 1 hàm chức năng, vd: *Calculate()*  
- ***Step 3: Create a test file***: Viết hàm test-case, Unit-test cho *Calculate()*  
- ***Step 4: Running Tests***

## Table Driven Testing
Table Driven chính là các model data input&output mong muốn của func.

Nguyên lý của Table Driven chính là xây dựng range `input` (array model) để ***expected*** === ***result***  

```go
expected := fmt.Sprintf("Hello %s", name)
result := sayHello(name)
```

```go
func TestTableCalculate(t *testing.T) {
    var tests = []struct {
        input    int
        expected int
    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    for _, test := range tests {
        if output := Calculate(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}

```
## Test Coverage
Thường thì Dev không viết hết được các case input data (in-range, out-range) (Table Driven) nên Go hỗ trợ test coverage
> Test coverage is measurement percentage of code coverage in your application.

### Một số lệnh coverage

> $ go test -cover

> $ go test -coverprofile=export_cover_out_files

> $ go tool cover -html=cover_out -o cover_out.html

## Benchmarks with Go
***Benchmarks*** là một công cụ nằm trong Test của Go, nhưng sẽ tách thành một bài viết riêng.

With Benchmarking you can measure performance of the code and see the impact of the changes you make the code so you can optimized your source code.

## Documenting Go Code
***Documenting Go Code With Example***

Như đã nói ở phần trước
> unit test may also act as documentation for new developers on the project.

Add Example (docs) approach for sayHello function like following.
```go
func ExamplesayHello() {
	fmt.Println(sayHello("Mert"))
	// Output: Hello Mert
}
```
- Example approach must start with ***Example*** prefix and following by existing function name
- fmt package is import to list what you expected and match the output
- Output: is document the expected output

Run the following command

> $ go test -v

## Ref
- [2: Unit Testing in Go (Golang)](https://medium.com/@snassr/unit-testing-in-go-golang-3a856ee3b9ba)
- [3: How to Write Unit Test in Go](https://medium.com/yemeksepeti-teknoloji/how-to-write-unit-test-in-go-1df2b98ad510)
- [4: writing unit test in golang easily](https://towardsdev.com/writing-unit-test-in-golang-easily-5fee03c653bb)
- [9: Testing in golang](https://medium.com/@jhyoocoderusa/testing-in-golang-ab89930a40f6) -Sơ sài
- [5 Testing Tips in Go](https://medium0.com/star-gazers/5-testing-tips-in-go-3b7f79a546da)
- [How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)
- [Unit Testing CLI Programs in Go](https://medium.com/swlh/unit-testing-cli-programs-in-go-6275c85af2e7)
- [Advanced Testing Tips in Go](https://medium0.com/@mert-akkaya/advanced-testing-tips-in-go-1b4d8eec82a0)
- [Implementing testing in Golang](https://dev.to/lucasnevespereira/implementing-testing-in-golang-4mcp)
- [Unit Testing in Go](https://www.pullrequest.com/blog/unit-testing-in-go/)
- [An Introduction to Testing in Go](https://tutorialedge.net/golang/intro-testing-in-go/)
- [Viết Unit Test trong project với Golang](https://viblo.asia/p/viet-unit-test-trong-project-voi-golang-bWrZnXv95xw)
- [Một số framework dùng cho viết Unit Test cho Golang](https://www.tma.vn/Hoi-dap/Cam-nang-nghe-nghiep/Golang-va-Unit-test/25459)
- [Step by step How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)
- [Golang — basic unit testing and benchmarking](https://sherryalex29.medium.com/golang-basic-unit-testing-and-benchmarking-10ce064f1081)

