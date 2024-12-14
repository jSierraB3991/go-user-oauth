package gooauthservice

import (
	"context"

	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) GetUserByUserId(ctx context.Context, keycloakId string) (*gooauthrest.User, error) {
	return nil, nil
}

func (s *GoOauthService) GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error) {
	return nil, nil
}
