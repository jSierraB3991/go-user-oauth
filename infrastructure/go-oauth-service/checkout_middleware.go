package gooauthservice

import (
	"net/http"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) CheckoutMiddleware(requets *http.Request) bool {
	path := gooauthmapper.ConvertPathToRegex(requets.URL.Path)
	allow := eliotlibs.PublicMiddleWare(path, requets.Method)
	if allow {
		return true
	}
	return false
}
