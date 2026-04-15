package gooauthservice

import (
	"log"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v3"
	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func GetCasbinConfig(configPath string) *casbin.Enforcer {
	ce, err := casbin.NewEnforcer(configPath+"/model.conf", configPath+"/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	return ce
}

func CasbinEchoConfig(requets *http.Request, e *casbin.Enforcer) error {

	path := gooauthmapper.ConvertPathToRegex(requets.URL.Path)
	method := requets.Method
	allow := eliotlibs.PublicMiddleWare(path, method)
	if allow {
		return nil
	}

	headers := requets.Header[gooauthlibs.HeaderAuthorization]
	if len(headers) <= 0 {
		return nil
	}

	if strings.TrimSpace(headers[0]) == gooauthlibs.ALONE_BEARER_HEADER {
		return nil
	}

	roleName, err := getHeaderJwtToken(requets, gooauthlibs.ROLE_NAME)
	if err != nil {
		log.Printf("Error GET ROLE NAME %v", err)
		return nil
	}
	// #nosec G706
	log.Printf(
		"roleName=%q path=%q method=%q",
		eliotlibs.SanitizeLog(roleName),
		eliotlibs.SanitizeLog(path),
		eliotlibs.SanitizeLog(method),
	)
	ok, err := e.Enforce(roleName, path, method)
	if err != nil || !ok {
		log.Println(err)
		return gooautherror.InvalidCasbinAccess{}
	}

	return nil
}

func getHeaderJwtToken(requet *http.Request, header string) (string, error) {
	stringInterface, err := gooauthlibs.GetClaimByToken(requet.Header[gooauthlibs.HeaderAuthorization][0], header)
	if err != nil {
		return "", err
	}

	if stringInterface != nil {
		return stringInterface.(string), nil
	}

	return "", nil
}
