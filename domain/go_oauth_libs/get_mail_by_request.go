package gooauthlibs

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func GetClaimByToken(tokenString, claim string) (interface{}, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jsierralibs.RemovePrefixInString(tokenString, BEARER_HEADER_PREFIX), jwt.MapClaims{})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims[claim], nil
	}
	return nil, errors.New("INVALID_CLAIMS")
}

func GetHeaderMail(requet *http.Request) (string, error) {
	var email string
	emailInterface, err := GetClaimByToken(requet.Header[HeaderAuthorization][0], EMAIL_CONSTANT)
	if err != nil {
		return "", err
	}

	if emailInterface != nil {
		return emailInterface.(string), nil
	}

	return email, nil
}
