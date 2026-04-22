package gossert

import "testing"

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
