## Api library for Market Steam on Go (golang).
## Installation

To install marketapi, simply run:
```
$ go get github.com/soluchok/marketapi
```
## Example
```go
package main

import (
    "fmt"
    "github.com/soluchok/marketapi"
)

func main() {
    dota, err := marketapi.NewDota2API("O3HJPmn98Q3Gkf0ujED3ZH479w625Zy")
    if err != nil {
        fmt.Println(err)
    }
    resp, err := dota.PingPong()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(resp.Success)
    fmt.Println(resp.Ping)
}
// {"success":true,"ping":"pong"}
// true
// pong
```
