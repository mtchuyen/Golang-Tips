# Benchmark


## Syntax

### Syntax for File

Benchmarks are placed inside ***_test.go*** files

### Syntax for Func

```func BenchmarkXyy (b *testing.B)```

- `Benchmark` is prefix
- `X` is upper-case char

### Syntax for Command

```go test -bench=<BenchFunc>```

Ex: `go test -bench=Xyy`

#### 1. Run all func

- `go test -bench=.` : run all func in this file (Test and Benchmark)

- `go test -run=AAA -bench=.`: run all func with 1 kind/mode in this file

`AAA` is mode: `Test` or `Bench`

#### 2. Run only one

Ex: `go test -run=Bench -bench=GenGUIDbyTime`

### Syntax for Output
```
goos: linux
goarch: amd64
pkg: first_Go/logtetbench
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkGenGUIDbyTime-6   	 2892226	       405.6 ns/op
PASS
ok  	first_Go/logtetbench	1.597s
```

#### BenchmarkGenGUIDbyTime-6
- `Benchmark` is prefix
- `GenGUIDbyTime` is name func
- ***-6*** 
#### 2892226
`2892226`: the number of operation in this bench's run

#### 405.6 ns/op
`405.6 ns/op`: time to run 1 operation

#### ok  	first_Go/logtetbench	1.597s
Total time to excutive this bench's run.
