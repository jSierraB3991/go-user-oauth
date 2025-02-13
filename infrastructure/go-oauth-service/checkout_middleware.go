package gooauthservice

import (
	"log"
	"net/http"
	"strings"

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

	if strings.Trim(headers[0], " ") == gooauthlibs.ALONE_BEARER_HEADER {
		return false
	}

	err := s.repo.SavePath(requets.URL.Path, requets.Method)
	if err != nil {
		log.Println(err)
	}
	return true
}
