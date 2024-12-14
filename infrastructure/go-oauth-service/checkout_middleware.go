package gooauthservice

import (
	"net/http"

	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) CheckoutMiddleware(requets *http.Request) bool {
	allow := jsierralibs.PublicMiddleWare(requets.URL.Path, requets.Method)
	if allow {
		return true
	}

	headers := requets.Header[gooauthlibs.HeaderAuthorization]
	if len(headers) <= 0 {
		return false
	}

	//token := headers[0]
	return true
}
