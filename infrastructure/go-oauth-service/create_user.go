package gooauthservice

import (
	"context"
	"strconv"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
)

func (s *GoOauthService) CreateUser(ctx context.Context, userParam gooauthrequest.CreateUser, role string, attributes *map[string][]string) (string, error) {
	userParam.UserName = userParam.Email
	roleUser, err := s.repo.GetRoleByName(ctx, role)
	if err != nil {
		return "", err
	}

	encryptPasword, err := s.passwordService.EncryptPassword(userParam.Password)
	if err != nil {
		return "", err
	}
	user := gooauthmapper.GetUserByCreate(userParam, roleUser, encryptPasword)

	err = s.repo.SaveUser(ctx, user)
	if err != nil {
		return "", err
	}

	err = s.repo.SaveAttributtes(ctx, user.UserId, gooauthmapper.GetAttributtes(attributes))
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(user.UserId)), nil
}
