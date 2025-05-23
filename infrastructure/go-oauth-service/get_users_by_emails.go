package gooauthservice

import (
	"context"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) GetUsersByEmail(ctx context.Context, emails []string) ([]gooauthrest.User, error) {
	usersData, err := s.repo.GetUsersByEmail(ctx, emails)
	if err != nil {
		return nil, err
	}
	return gooauthmapper.GetUsersRest(usersData), nil
}
