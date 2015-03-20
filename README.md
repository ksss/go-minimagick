go-minimagick
===

go-minimagick is a command wrapper for imagemagick inspired by [minimagick/minimagick](https://github.com/minimagick/minimagick)

```go
package main

import (
        "fmt"

        "github.com/ksss/go-minimagick/convert"
        "github.com/ksss/go-minimagick/info"
)

func main () {
        shellOutBytes, err = convert.New().Resize("100x100").Add("some.jpg", "out.jpg").Run()

        width, err := info.Width("some.jpg")
        height, err := info.Height("some.jpg")
        fmt.Println(width, height)
}
```
