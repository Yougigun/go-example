# Golang test common usage

## Test
```bash
go test . 
go test file1 file2 file3
go test -timeout 30s -run ^TestGetAppInfoOne$ themis/frontendsrv/dao/mysql
```
## Coverage
```bash
$ go test -cover .
ok      calc        0.001s  coverage: 100.0% of statements

$ go test ./... -coverprofile cover.out
ok      calc        0.001s  coverage: 100.0% of statements
ok      calc/format 0.001s  coverage: 100.0% of statements

$ go tool cover -func cover.out
calc/calc.go:4:              Add             100.0%
calc/calc.go:9:              Subtract        100.0%
calc/calc.go:14:             Multiply        100.0%
calc/format/format.go:8:     Print           100.0%
total:                       (statements)    100.0%

$ go tool cover -func cover.out | grep total:
total:                       (statements)    100.0%
```
## html
``` bash
go tool cover -html=coverage.out
```
## CI
$ go tool cover -func cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}' | xargs -I {} curl \
    --header "Authorization: Token MY_TOKEN" \
    --data value="{}" \
    --data sha="MY_SHA" \
    https://seriesci.com/api/repos/:owner/:repo/:series/values

# Benchmark
## Basic
```bash
$ go test -v -bench=. benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Add-4           20000000         0.33 ns/op
PASS
ok          command-line-arguments        0.700s
```
代码说明如下：
第 1 行的-bench=.表示运行 benchmark_test.go 文件里的所有基准测试，和单元测试中的-run类似。
第 4 行中显示基准测试名称，2000000000 表示测试的次数，也就是 testing.B 结构中提供给程序使用的 N。“0.33 ns/op”表示每一个操作耗费多少时间（纳秒）。

## Test Time
```bash
$ go test -v -bench=. -benchtime=5s benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Add-4           10000000000                 0.33 ns/op
PASS
ok          command-line-arguments        3.380s
```
## Memory
```go
func Benchmark_Alloc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("%d", i)
    }
}
```
在命令行中添加-benchmem参数以显示内存分配情况，参见下面的指令：
```bash
$ go test -v -bench=Alloc -benchmem benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Alloc-4 20000000 109 ns/op 16 B/op 2 allocs/op
PASS
ok          command-line-arguments        2.311s
```
代码说明如下：
第 1 行的代码中-bench后添加了 Alloc，指定只测试 Benchmark_Alloc() 函数。
第 4 行代码的“16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。

## Time Control
```go
func Benchmark_Add_TimerControl(b *testing.B) {
    // 重置计时器
    b.ResetTimer()
    // 停止计时器
    b.StopTimer()
    // 开始计时器
    b.StartTimer()
    var n int
    for i := 0; i < b.N; i++ {
        n++
    }
}
```