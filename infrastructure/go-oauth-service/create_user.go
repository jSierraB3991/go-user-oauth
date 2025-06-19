package gooauthservice

import (
	"context"
	"strconv"
	"strings"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) CreateUser(ctx context.Context, userParam gooauthrequest.CreateUser, role string, attributes *map[string][]string) (string, error) {

	userEmail := strings.ToLower(eliotlibs.RemoveSpace(userParam.Email))
	userParam.Email = userEmail
	userParam.UserName = userEmail
	userParam.FirstName = eliotlibs.CapitalizeName(userParam.FirstName)
	userParam.LastName = eliotlibs.CapitalizeName(userParam.LastName)

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
