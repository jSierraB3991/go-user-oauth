package gooauthservice

import (
	"context"

	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
)

func (s *GoOauthService) ExistsUserAdministrator(ctx context.Context) (bool, error) {
	return s.repo.ExistsUserAdministrator(ctx)
}

func (s *GoOauthService) CreateUserAdministrator(ctx context.Context, userEmail, userpassword, appName string, attributes *map[string][]string) (string, error) {

	existsUserAdministrator, err := s.ExistsUserAdministrator(ctx)
	if err != nil {
		return "", err
	}

	if existsUserAdministrator {
		return "", nil
	}

	kUser := gooauthrequest.CreateUser{
		Email:       userEmail,
		UserName:    userEmail,
		Password:    userpassword,
		FirstName:   "Admin",
		LastName:    appName,
		Emailverify: true,
	}

	userId, err := s.CreateUser(ctx, kUser, gooauthlibs.ROLE_ADMIN, attributes)
	if err != nil {
		return "", err
	}

	return userId, nil
}
