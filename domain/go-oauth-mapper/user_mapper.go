package gooauthmapper

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
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

func GetUserRestAnAttributtes(user *gooauthmodel.User, attrs []gooauthmodel.UserAttributtes) *gooauthrest.User {
	var attrRes *map[string][]string
	for _, attr := range attrs {
		addToMap(attrRes, attr.NameAttributte, attr.ValueAttributtes)
	}
	return &gooauthrest.User{
		Email:      user.Email,
		Name:       user.Name,
		SubName:    user.SubName,
		Enabled:    user.Enabled,
		Role:       user.Role.RoleName,
		Attributes: attrRes,
	}
}

func addToMap(attrRes *map[string][]string, key string, value string) {
	// Si la clave no existe en el mapa, inicializamos su slice
	if _, exists := (*attrRes)[key]; !exists {
		(*attrRes)[key] = []string{}
	}

	// Agregar el valor al slice correspondiente
	(*attrRes)[key] = append((*attrRes)[key], value)
}
