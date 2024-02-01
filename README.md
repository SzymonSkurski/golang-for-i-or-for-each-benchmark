# Benchmarking

benchmark simple run `go test -bench=. -run=x`  
benchmark with set time consumption `go test -bench=. -run=x -benchtime=5s`  
benchmark with memory usage `go test -bench=. -run="x" -benchmem`  
benchmark with memory usage and store profile `go test -bench=. -run="x" -benchmem -memprofile mem.prof`  
benchmark with memory usage and store profiles `go test -bench=. -run="x" -benchmem -memprofile mem.prof -cpuprofile cpu.prof`  

# Profiling

pprof cpu profile `go tool pprof cpu.prof` or `go tool pprof mem.prof`
serve cpu.prof file on localhost:4001 `pprof -http=localhost:4001 cpu.prof`  
serve mem.prof file on localhost:4001 `pprof -http=localhost:4002 mem.prof`
