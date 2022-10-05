# Go Utility: `password`

## Installation

```bash
go get github.com/go-ginseng/password
```

## Hash Password

### hash password

```go
// import "github.com/go-ginseng/password"

h := password.Hash("password")
```

### verify password

```go
// import "github.com/go-ginseng/password"

h := password.Hash("password")
ok := password.Verify("password", h)
```
