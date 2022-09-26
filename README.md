# Go Utility: `password`

## Installation

```bash
go get github.com/nelsonlai-go/password
```

## Hash Password

### hash password

```go
// import "github.com/nelsonlai-go/password"

h := password.Hash("password")
```

### verify password

```go
// import "github.com/nelsonlai-go/password"

h := password.Hash("password")
ok := password.Verify("password", h)
```
