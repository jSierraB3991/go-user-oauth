package gooauthmapper

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func GetUserByCreate(userParam gooauthrequest.CreateUser, role *gooauthmodel.GoUserRole, password string) *gooauthmodel.GoUserUser {
	return &gooauthmodel.GoUserUser{
		Email:      userParam.Email,
		Name:       userParam.FirstName,
		SubName:    userParam.LastName,
		Password:   password,
		GoUserRole: *role,
		Enabled:    userParam.Emailverify,
	}
}

func GetAttributtes(attributes *map[string][]string) []gooauthmodel.GoUserUserAttributtes {
	if attributes == nil {
		return nil
	}
	var result []gooauthmodel.GoUserUserAttributtes
	for key, values := range *attributes {
		result = append(result, gooauthmodel.GoUserUserAttributtes{
			NameAttributte:   key,
			ValueAttributtes: values[0],
		})
	}
	return result
}
func GetUserRestAnAttributtes(user *gooauthmodel.GoUserUser, attrs []gooauthmodel.GoUserUserAttributtes) *gooauthrest.User {
	attrRes := make(map[string][]string)

	for _, attr := range attrs {
		addToMap(&attrRes, attr.NameAttributte, attr.ValueAttributtes)
	}

	return &gooauthrest.User{
		Id:         jsierralibs.GetFloatStringToUInt(user.UserId),
		Email:      user.Email,
		Name:       user.Name,
		SubName:    user.SubName,
		Enabled:    user.Enabled,
		Role:       user.GoUserRole.RoleName,
		Attributes: &attrRes,
	}
}

func GetUsersRest(users []gooauthmodel.GoUserUser) []gooauthrest.User {
	var result []gooauthrest.User
	for _, v := range users {
		data := GetUserRestAnAttributtes(&v, nil)
		result = append(result, *data)
	}
	return result
}

func addToMap(attrRes *map[string][]string, key string, value string) {
	// Si la clave no existe en el mapa, inicializamos su slice
	if _, exists := (*attrRes)[key]; !exists {
		(*attrRes)[key] = []string{}
	}

	// Agregar el valor al slice correspondiente
	(*attrRes)[key] = append((*attrRes)[key], value)
}

func GetAttributteUpdate(attrData []gooauthmodel.GoUserUserAttributtes,
	attrDataNews []gooauthmodel.GoUserUserAttributtes,
	userId uint) []gooauthmodel.GoUserUserAttributtes {

	var result []gooauthmodel.GoUserUserAttributtes

	for _, j := range attrDataNews {
		isSavePre := false
		for _, v := range attrData {
			if j.NameAttributte == v.NameAttributte {
				v.ValueAttributtes = j.ValueAttributtes
				result = append(result, v)
				isSavePre = true
				break
			}
		}

		if !isSavePre {
			result = append(result, gooauthmodel.GoUserUserAttributtes{
				NameAttributte:   j.NameAttributte,
				GoUserUserId:     userId,
				ValueAttributtes: j.ValueAttributtes,
			})
		}
	}

	for _, v := range attrData {
		isForSave := false
		for _, j := range attrDataNews {
			if v.NameAttributte == j.NameAttributte {
				isForSave = true
			}
		}

		if !isForSave {
			result = append(result, v)
		}
	}

	return result
}
