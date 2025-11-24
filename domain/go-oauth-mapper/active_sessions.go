package gooauthmapper

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func MapLoginSessionsToRest(sessions []gooauthmodel.GoUserDataLogin) []gooauthrest.LoginSessionRest {
	var result []gooauthrest.LoginSessionRest
	for _, session := range sessions {
		result = append(result, gooauthrest.LoginSessionRest{
			Id:                  session.UserDataLoginId,
			UserAgent:           session.UserAgent,
			Ip:                  session.Ip,
			IpResponse:          session.IpResponse,
			IsLoginWithPassword: session.IsLoginWithPassword,
			IsAvailable:         session.IsAvailable,
			Fecha:               session.Fecha,
		})
	}
	return result
}
