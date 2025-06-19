package gooauthservice

import (
	"context"
	"strings"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) GetUserByEmail(ctx context.Context, email string) (*gooauthrest.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, strings.ToLower(email))
	if err != nil {
		return nil, err
	}

	attributtes, err := s.repo.GetAttributtesByUserId(ctx, user.UserId)
	if err != nil {
		return nil, err
	}

	return gooauthmapper.GetUserRestAnAttributtes(user, attributtes), nil
}
