package gooauthservice

import (
	"context"

	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) GetUserByRole(ctx context.Context, role string) ([]*gooauthrest.User, error) {
	return nil, nil
}
