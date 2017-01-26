package errorx

import (
	"fmt"
	"log"
	"net/http"
)

// Err Represents an http error
type Err struct {
	StdError error
	UserMsg  string
	ErrCode  int
}

// New returns a new Err
func New(usermsg string, err error, code int) *Err {
	return &Err{UserMsg: usermsg, StdError: err, ErrCode: code}
}

// Errorf will return a new Err with the user message formatted
// ErrCode will is set to InternalServerError
// StdError is set to the user message
func Errorf(format string, a ...interface{}) *Err {
	return New(fmt.Sprintf(format, a...), fmt.Errorf(format, a...), InternalServerError)
}

// WithError will return a new Err with the user message and StdError
// set to err.Error()
// Code is set to InternalServerError
func WithError(err error) *Err {
	return New(err.Error(), err, InternalServerError)
}

// Error satisfies the error interface
func (e *Err) Error() string {
	return fmt.Sprintf("%v %v %v", e.UserMsg, e.StdError, e.ErrCode)
}

// Code will add a error code to the error
func (e *Err) Code(code int) *Err {
	e.ErrCode = code
	return e
}

// Msg will add a user message to the error
func (e *Err) Msg(msg string) *Err {
	e.UserMsg = msg
	return e
}

// Err will wrap a standard error
func (e *Err) Err(err error) *Err {
	e.StdError = err
	return e
}

// NewIfErr returns a new Err if orginal error is not nil
// otherwise returns nil
func NewIfErr(msg string, err error, code int) *Err {
	if err == nil {
		return nil
	}
	return New(msg, err, code)
}

// Write sends the user message and http error code to the user
// and logs the original error
func (e *Err) Write(w http.ResponseWriter, r *http.Request) {
	code := e.ErrCode
	txt, ok := errorCodes[code]
	if !ok {
		code = 500
		txt = errorCodes[code]
	}
	http.Error(w, fmt.Sprintf("%v: %v", txt, e.UserMsg), code)
	log.Printf("[ERROR] %v %v %v %v %v %v %v", txt, code, e.StdError, e.UserMsg, r.RemoteAddr, r.Method, r.URL)
}

// WriteErr will check if the err is a errorx and write if it is
// if not it will create an errorx using WithError and write it
func WriteErr(w http.ResponseWriter, r *http.Request, err error) {
	if e, ok := err.(*Err); ok {
		e.Write(w, r)
		return
	}
	WithError(err).Write(w, r)
}
