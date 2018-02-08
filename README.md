# log4e

An [echo](https://github.com/labstack/echo) logger implementation in logrus

## Example

```go
package main

import (
    "github.com/labstack/echo"
    log "github.com/kevinclcn/log4e"
)

func main() {
    e := echo.New()
    e.Logger = log.Logger()
    log.Info("starting server")
    e.Logger.Fatal(e.Start(":1213"))
}

```
