# Running the benchmarks

``` shell
~/GolangUnittesting$ cd benchmark/
~/GolangUnittesting/benchmark$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: example
BenchmarkConcatStrings-8              20000000    58.9 ns/op
BenchmarkConcatStringsWithAllocations-8    20000000    58.9 ns/op    8 B/op    1 allocs/op
BenchmarkSub/ShortStrings-8           50000000    30.2 ns/op
BenchmarkSub/LongStrings-8            10000000   120.3 ns/op
PASS
ok  	example	2.355s
```
> BenchmarkConcatStrings-8: Name of the benchmark and number of CPU cores used.<br>
> 20000000: Number of iterations.<br>
> 58.9 ns/op: Average time per operation.<br>
> 8 B/op: Bytes allocated per operation (if -benchmem flag is used).<br>
> 1 allocs/op: Number of allocations per operation (if -benchmem flag is used).<br>

