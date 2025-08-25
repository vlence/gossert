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

// An assertion error. As the name implies this error type is used
// to denote an assertion failing. The purpose of this type is to
// help distinguish assertion errors from other error types when
// recovering from a panic.
type AssertionError struct {
        msg string
}

func (err *AssertionError) Error() string {
        return err.msg
}

// Ok panics if the given condition is not true. msg is the error message
// used when Ok panics.
func Ok(cond bool, msg string) {
        if !cond {
                panic(&AssertionError{msg})
        }
}
