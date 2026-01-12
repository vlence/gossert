package gossert

import "testing"

func TestOkDoesNotPanicWhenGivenTrue(t *testing.T) {
        defer func () {
                if r := recover(); r != nil {
                        t.Errorf("Ok panicked; expected Ok to not panic")
                }
        }()
        
        errmsg := "error message"
        Ok(true, errmsg)
}

func TestOkPanicsWhenGivenFalse(t *testing.T) {
        errmsg := "error message"
        
        defer func (errmsg string) {
                r, ok := recover().(*Error)
                
                if !ok {
                        t.Fatalf("could not recover error; got %v", r)
                }

                if r.msg != errmsg {
                        t.Fatalf("error message did not match panic message; got \"%s\"", r.Error())
                }
        }(errmsg)
        
        Ok(false, errmsg)

        t.Errorf("expected Ok to panic")
}
