package gooauthmapper

import (
	"encoding/json"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func MapLoginSessionsToRest(sessions []gooauthmodel.GoUserDataLogin, aesEncrypt, tokenString string) []gooauthrest.LoginSessionRest {
	var result []gooauthrest.LoginSessionRest
	for _, session := range sessions {

		ipDecrypt, err := eliotlibs.Decrypt(session.Ip, aesEncrypt)
		if err != nil {
			ipDecrypt = "Desconocido"
		}
		tokenDecrypt, err := eliotlibs.Decrypt(session.Token, aesEncrypt)
		isThis := false
		if err == nil {
			isThis = tokenDecrypt == tokenString
		}

		var ipResponse gooauthrequest.IPInfoRequest
		if session.IpResponse != "" {
			err := json.Unmarshal([]byte(session.IpResponse), &ipResponse)
			if err != nil {
				ipResponse = gooauthrequest.IPInfoRequest{Country: "Desconocido", City: "Desconocido"}
			}
		}
		result = append(result, gooauthrest.LoginSessionRest{
			Id:                  session.UserDataLoginId,
			UserAgent:           session.UserAgent,
			Ip:                  ipDecrypt,
			IpResponse:          ipResponse,
			IsLoginWithPassword: session.IsLoginWithPassword,
			IsAvailable:         session.IsAvailable,
			Fecha:               session.Fecha,
			IsThisSession:       isThis,
		})
	}
	return result
}
