# Kutex [![GoDoc](https://godoc.org/github.com/4vn/kutex?status.svg)](https://godoc.org/github.com/4vn/kutex)

Kutex = key + mutex

### Import

``` go
import "github.com/4vn/kutex"
```

### Usage

``` go
ku := kutex.New(1000)
userId := "gaia"
ku.Lock(userId)
// Do something
ku.Unlock(userId)
```

### License

MIT
