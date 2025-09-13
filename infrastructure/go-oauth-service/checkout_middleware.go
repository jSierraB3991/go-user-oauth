package gooauthservice

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) CheckoutMiddleware(requets *http.Request) bool {

	ctx := requets.Context()
	path := gooauthmapper.ConvertPathToRegex(requets.URL.Path)
	allow := eliotlibs.PublicMiddleWare(path, requets.Method)
	if allow {
		return true
	}

	pathId, err := s.repo.SavePath(ctx, path, requets.Method)
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
		err = s.repo.SavePathRole(ctx, pathId, roleName)
		if err != nil {
			log.Printf("Error save role path %v", err)
		}
	}

	if strings.Contains(path, gooauthlibs.ADMIN_ROUTES) {
		emailUser, err := GetHeaderJwtToken(requets, "email")
		if err != nil {
			log.Printf("Error GET EMAIL %v", err)
			return false
		}

		userData, err := s.repo.GetUserByEmail(ctx, emailUser)
		if err != nil {
			log.Printf("Error getting user by email %v", err)
			return false
		}

		if userData == nil {
			log.Printf("User not found with email %s", emailUser)
			return false
		}
		if userData.GoUserRole.RoleName != gooauthlibs.ROLE_ADMIN {
			log.Printf("User %s with role %s tried to access admin route %s", userData.Name+" "+userData.SubName, roleName, path)
			return false
		}
	}
	bodyBytes, err := io.ReadAll(requets.Body)
	if err == nil {
		// Restaurar el body para que se pueda usar de nuevo
		requets.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		log.Println("Body recibido:", string(string(bodyBytes)))
	}

	return false
}

func GetHeaderJwtToken(requet *http.Request, header string) (string, error) {
	stringInterface, err := gooauthlibs.GetClaimByToken(requet.Header[gooauthlibs.HeaderAuthorization][0], header)
	if err != nil {
		return "", err
	}

	if stringInterface != nil {
		return stringInterface.(string), nil
	}

	return "", nil
}
