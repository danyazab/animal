# Animals

Api documentation available in [Swagger](http://localhost:8000/swagger)

## Build & Run (Locally)
### Prerequisite
- go 1.18

```bash
 // run build
 root$ go build -v cmd/api.go
```

## CodeStyle & UnitTest
```bash
// fix code style
go fmt ./...

// run code analyzer
 golangci-lint run ./...

// run unit-test
go test -v -race -timeout 30s ./...
```

