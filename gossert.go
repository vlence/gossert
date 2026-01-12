/*
Package gossert exports functions to perform assertions. Assertions are checks
that your program must pass. If an assertion fails then the current goroutine
will panic.

Usage:

	package main

	import "github.com/vlence/gossert"

	func main() {
	        gossert.Ok(true, "you will never see this error message")
	}
*/
package gossert

// Set this to false if you don't want to panic when an assertion fails.
// This is set to true by default. Some may want to turn off panics in
// production but I recommend leaving it on because if your assertion
// fails it means your program is no longer in a state that it expects;
// executing from this point on would be undefined behaviour.
var PanicOnAssertionError = true

// An assertion error. The purpose of this type is to make it simple
// to determine if an error is an assertion error.
type Error struct {
	msg   string
	cause error
}

// Error returns the assertion error message as "assert: msg"
func (e *Error) Error() string {
	if e.cause == nil {
		return "assert: " + string(e.msg)
	} else {
		return "assert: " + string(e.msg) + "\n\ncause: " + e.cause.Error()
	}
}

// Cause returns the error that led to this assertion error.
func (e *Error) Cause() error {
	return e.cause
}

// Ok panics if the given condition is not true. msg is the error message
// used when Ok panics.
func Ok(cond bool, msg string) {
	if !PanicOnAssertionError {
		return
	}

	if !cond {
		panic(&Error{msg, nil})
	}
}

// NilErr panics with message msg if err is not nil. The
// cause of the assertion error will be err.
func NilErr(err error, msg string) {
	if !PanicOnAssertionError {
		return
	}

	if err != nil {
		panic(&Error{msg, err})
	}
}
