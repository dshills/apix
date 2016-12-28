package route

import (
	"net/http"
	"strings"

	"github.com/bouk/httprouter"
	"github.com/dshills/apix/ierror"
	"github.com/dshills/apix/token"
)

// Router is a wrapper for httprouter.Router
type Router struct {
	httprouter.Router
}

// New returns a new Router
func New() *Router {
	r := httprouter.New()
	return &Router{*r}
}

// Secured sets a secured route using jwt middleware
func (r *Router) Secured(method, path string, hf http.HandlerFunc) {
	m := strings.ToUpper(method)
	switch m {
	case "GET":
		r.GET(path, WrapJWT(hf))
	case "POST":
		r.POST(path, WrapJWT(hf))
	case "DELETE":
		r.DELETE(path, WrapJWT(hf))
	case "PUT":
		r.PUT(path, WrapJWT(hf))
	}
}

// Insecured sets a route
func (r *Router) Insecured(method, path string, hf http.HandlerFunc) {
	m := strings.ToUpper(method)
	switch m {
	case "GET":
		r.GET(path, hf)
	case "POST":
		r.POST(path, hf)
	case "DELETE":
		r.DELETE(path, hf)
	case "PUT":
		r.PUT(path, hf)
	}
}

// WrapJWT wraps a HandlerFunc with JWT security
func WrapJWT(hf http.HandlerFunc) http.HandlerFunc {
	return WrapHandler(token.Handler(http.HandlerFunc(hf)))
}

// WrapHandler wraps a Handler and returns a HandlerFunc
func WrapHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

// NotFound handles a 404 Not Found request
func NotFound(w http.ResponseWriter, r *http.Request) {
	ierr := ierror.New("", nil, 404)
	ierr.Write(w, r)
}
