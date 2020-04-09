# Kutex - key mutex

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
