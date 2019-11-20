## Run Benchmarks

`go test -v bench=.`

## Find the data race

`go test -race -v bench=.`

… or build with `go build -race` and then run it.
