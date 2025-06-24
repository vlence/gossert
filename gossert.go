/*
Package gossert exports functions to perform assertions. Assertions are checks
that your program must pass. If an assertion fails then the current goroutine
will panic.

Usage:

        package main

        import "github.com/vlence/gossert"

        func main() {
                gossert.Ok(true, 'you will never see this error message') 
        }
*/
package gossert

import "errors"

// Ok panics if the given condition is not true. msg is the error message
// used when Ok panics.
func Ok(cond bool, msg string) {
        if !cond {
                panic(errors.New(msg))
        }
}
