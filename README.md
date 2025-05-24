# go-serror

serror is errors with sources

# Usage

```go
package main

import (
  "context"
  "fmt"
  "os"

  "github.com/skaji/go-serror"
)

func doSomething(ctx context.Context) error {
  if ctx == nil {
    return serror.Error("context must not be nil")
  }
  // do something...
  if err != nil {
    return serror.Wrap(err)
  }
  return nil
}

func main() {
  if err := doSomething(nil); err != nil {
    fmt.Println(serror.Wrap(err))
    // context must not be nil
    //   at main.go:13
    //   at main.go:24
    os.Exit(1)
  }
}
```

# Author

Shoichi Kaji

# License

MIT
