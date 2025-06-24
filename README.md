# gossert

A simple assertion library written in Go. Gossert exports functions that can be used to perform
assertions. Assertions are checks that your program must pass. If an assertion fails the current
goroutine panics.

## Usage

```go
package main

import "github.com/vlence/gossert"

func main() {
    gossert.Ok(true, "this will not panic")
    gossert.Ok(false, "this will panic")
}
```
