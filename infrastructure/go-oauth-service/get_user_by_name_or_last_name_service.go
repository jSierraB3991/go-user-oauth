package gooauthservice

import (
	"context"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) GetUserByName(ctx context.Context, name string, page *eliotlibs.Paggination) ([]string, error) {

	usersLikeName, err := s.repo.GetUsersByNamePage(ctx, page, name)
	if err != nil {
		return nil, err
	}
	return GetEmailByUserModel(usersLikeName), nil
}

func GetEmailByUserModel(usersLikeName []gooauthmodel.GoUserUser) []string {
	var result []string
	for _, v := range usersLikeName {
		result = append(result, v.Email)
	}
	return result
}
