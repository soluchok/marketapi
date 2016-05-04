## Api library for Market Steam on Go (golang).

## Example
```go
package main

import (
    "fmt"
    "marketapi"
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
