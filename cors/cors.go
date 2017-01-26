package cors

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// AccessControl is a CORS implementation
type AccessControl struct {
	AllowCredentials bool
	AllowHeaders     []string
	AllowMethods     []string
	AllowOrigin      []string
	ExposeHeaders    []string
	MaxAge           time.Duration
}

var globalDefault *AccessControl

// SetDefault will set a global default Access Control response
func SetDefault(ac *AccessControl) {
	globalDefault = ac
}

func mergeDefaults(ac *AccessControl) *AccessControl {
	switch {
	case globalDefault == nil && ac == nil:
		return &AccessControl{}
	case globalDefault == nil:
		return ac
	case ac == nil:
		return globalDefault
	}
	newAC := *globalDefault
	if ac.AllowCredentials {
		newAC.AllowCredentials = ac.AllowCredentials
	}
	if len(ac.AllowHeaders) > 0 {
		newAC.AllowHeaders = ac.AllowHeaders
	}
	if len(ac.AllowMethods) > 0 {
		newAC.AllowMethods = ac.AllowMethods
	}
	if len(ac.AllowOrigin) > 0 {
		newAC.AllowOrigin = ac.AllowOrigin
	}
	if len(ac.ExposeHeaders) > 0 {
		newAC.ExposeHeaders = ac.ExposeHeaders
	}
	if ac.MaxAge > 0 {
		newAC.MaxAge = ac.MaxAge
	}
	return &newAC
}

// Handler is a handler for CORS access
func Handler(ac *AccessControl) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		access := mergeDefaults(ac)
		if access.AllowCredentials {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		if len(access.AllowHeaders) > 0 {
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(access.AllowHeaders, ", "))
		}
		if len(access.AllowMethods) > 0 {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(access.AllowMethods, ", "))
		}
		if len(access.AllowOrigin) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", strings.Join(access.AllowOrigin, ", "))
		}
		if len(access.ExposeHeaders) > 0 {
			w.Header().Set("Access-Control-Expose-Headers", strings.Join(access.ExposeHeaders, ", "))
		}
		if access.MaxAge > 0 {
			w.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%v", access.MaxAge))
		}
		fmt.Fprint(w, "")
	})
}
