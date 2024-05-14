# Benchmarking

benchmark only one test `go test -bench="BenchmarkCloneForI"`  
benchmark simple run `go test -bench=.`  
benchmark with set time `go test -bench=. -benchtime=5s`
benchmark with set n amount `go test -bench=. -benchtime=100x`  
benchmark with set cpu (bench for each set cpu) `go test -bench=. -cpu=1,2,3,4,5,6`  
benchmark with memory usage `go test -bench=. -benchmem`  
benchmark with memory usage and store profile `go test -bench=. -benchmem -memprofile mem.prof`  
benchmark with memory usage, cpu usage and store profiles `go test -bench=. -run="x" -benchmem -memprofile mem.prof -cpuprofile cpu.prof`  

use `go help testflag` to see all possible options

# Profiling

pprof cpu profile `go tool pprof cpu.prof` or `go tool pprof mem.prof`
serve cpu.prof file on localhost:4001 `pprof -http=localhost:4001 cpu.prof`  
serve mem.prof file on localhost:4001 `pprof -http=localhost:4002 mem.prof`


# Stat
`benchstat result.prof another_result.prof` compare benchmark results
