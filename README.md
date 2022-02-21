# Example of Go WebAssembly

### Build

```
$ cd example
$ GOOS=js GOARCH=wasm go build -o main.wasm
```


### Execute

```
$ go run server.go
# http://localhost:8080/example/

```
