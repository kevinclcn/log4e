# log4e

logrus for echo framework

## Example

```go
package main

import (
    "github.com/labstack/echo"
    "github.com/kevinclcn/log4e"
)

func main() {
    e := echo.New()
    e.Logger = log4e.Logger()
    log4e.Info("starting server")
    e.Logger.Fatal(e.Start(":1213"))
}

```
