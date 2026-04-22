package gossert

import (
	"errors"
	"testing"
)

func TestOkDoesNotPanicWhenGivenTrue(t *testing.T) {
        defer func () {
                if r := recover(); r != nil {
                        t.Fatalf("expected to recover nil but got %#v", r)
                }
        }()
        
        msg := "error message"
        Ok(true, msg)
}

func TestOkPanicsWhenGivenFalse(t *testing.T) {
        msg := "error message"
        
        defer func (msg string) {
		r := recover()

		if r == nil {
			t.Fatalf("expected to recover assertion error but got nil")
		}

		err, ok := r.(*Error)
                
                if !ok {
                        t.Fatalf("could not recover error; got %#v", r)
                }

                if err.msg != msg {
                        t.Fatalf("error message did not match panic message; got \"%s\"", err.msg)
                }
        }(msg)
        
        Ok(false, msg)

        t.Errorf("Ok did not panic")
}

func TestNilErrPanicsWhenErrIsNotNil(t *testing.T) {
	cause := errors.New("assertion error cause")
	msg := "assertion failed"

	defer func(msg string, cause error) {
		r := recover()

		if r == nil {
			t.Fatalf("expected assertion error but got nil")
		}

		err, ok := r.(*Error)

		if !ok {
			t.Fatalf("expected to recover assertion error but got %#v", r)
		}

		if err == nil {
			t.Fatalf("expected assertion error but got nil")
		}

		if err.cause != cause {}

		if err.msg != msg {}
	}(msg, cause)

	NilErr(cause, msg)

	t.Errorf("NilError did not panic")
}

func TestNilErrDoesNotPanicWhenErrIsNil(t *testing.T) {
	defer func() {
		r := recover()

		if r != nil {
			t.Logf("expected nil but got %#v", r)
		}

		err, ok := r.(error)

		if ok {
			t.Fatalf("unexpected error %s", err.Error())
		}
	}()

	NilErr(nil, "I should not see this")
}
