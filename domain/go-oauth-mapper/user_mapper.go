package gooauthmapper

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
)

func GetUserByCreate(userParam gooauthrequest.CreateUser, role *gooauthmodel.Role, password string) *gooauthmodel.User {
	return &gooauthmodel.User{
		Email:    userParam.Email,
		Name:     userParam.FirstName,
		SubName:  userParam.LastName,
		Password: password,
		Role:     *role,
		Enabled:  userParam.Emailverify,
	}
}

func GetAttributtes(attributes *map[string][]string) []gooauthmodel.UserAttributtes {
	var result []gooauthmodel.UserAttributtes
	for key, values := range *attributes {
		result = append(result, gooauthmodel.UserAttributtes{
			NameAttributte:   key,
			ValueAttributtes: values[0],
		})
	}
	return result
}
