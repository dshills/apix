package ierror

import (
	"fmt"
	"log"
	"net/http"
)

// ErrCode types
type ErrCode int

// ErrCode constants
const (
	InternalError     ErrCode = 500
	NoRecordError             = 404
	NotFoundError             = 404
	ForbiddenError            = 403
	UnauthorizedError         = 401
	BadRequestError           = 400
	InvalidInputError         = 400
)

// Err Represents an http error
type Err struct {
	Org     error
	UserMsg string
	Code    ErrCode
}

func (e *Err) Error() string {
	return fmt.Sprintf("%v %v %v", e.Code, e.Org, e.UserMsg)
}

// New returns a new Err
func New(msg string, err error, code ErrCode) *Err {
	return &Err{Org: err, UserMsg: msg, Code: code}
}

// NewIfErr returns a new Err if orginal error is not nil
// otherwise returns nil
func NewIfErr(msg string, err error, code ErrCode) *Err {
	if err == nil {
		return nil
	}
	return New(msg, err, code)
}

func (e *Err) Write(w http.ResponseWriter, r *http.Request) {

	switch e.Code {
	case 401:
		http.Error(w, fmt.Sprintf("Unauthorized 401: %v\n", e.UserMsg), 401)
	case 403:
		http.Error(w, fmt.Sprintf("Forbidden 403: %v\n", e.UserMsg), 403)
	case 404:
		http.Error(w, fmt.Sprintf("Not Found 404: %v\n", e.UserMsg), 404)
	case 500:
		http.Error(w, fmt.Sprintf("Internal Server Error 500: %v\n", e.UserMsg), 500)
	default:
		http.Error(w, fmt.Sprintf("%v\n", e.UserMsg), int(e.Code))
	}

	log.Printf("[ERROR] Generated: %v User: %v %v %v %v", e.Org, e.UserMsg, r.RemoteAddr, r.Method, r.URL)
}
