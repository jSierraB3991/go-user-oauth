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

	pathId, err := s.repo.SavePath(requets.URL.Path, requets.Method)
	if err != nil {
		log.Printf("Error save path %v", err)
	}

	headers := requets.Header[gooauthlibs.HeaderAuthorization]
	if len(headers) <= 0 {
		return false
	}

	if strings.TrimSpace(headers[0]) == gooauthlibs.ALONE_BEARER_HEADER {
		return false
	}

	roleName, err := GetHeaderJwtToken(requets, "role_name")
	if err != nil {
		log.Printf("Error GET ROLE NAME %v", err)
		return false
	}

	if roleName != "" {
		err = s.repo.SavePathRole(pathId, roleName)
		if err != nil {
			log.Printf("Error save role path %v", err)
		}
	}
	return false
}

func GetHeaderJwtToken(requet *http.Request, header string) (string, error) {
	stringInterface, err := jsierralibs.GetClaimByToken(requet.Header[gooauthlibs.HeaderAuthorization][0], header)
	if err != nil {
		return "", err
	}

	if stringInterface != nil {
		return stringInterface.(string), nil
	}

	return "", nil
}
