package gooauthservice

import (
	"context"
	"strconv"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) GetUserByUserId(ctx context.Context, userId string) (*gooauthrest.User, error) {
	userInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUserById(uint(userInt))
	if err != nil {
		return nil, err
	}

	attributtes, err := s.repo.GetAttributtesByUserId(uint(userInt))
	if err != nil {
		return nil, err
	}

	return gooauthmapper.GetUserRestAnAttributtes(user, attributtes), nil
}

func (s *GoOauthService) GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error) {
	var result []gooauthrest.User
	for _, v := range keycloakUsersId {
		user, err := s.GetUserByUserId(ctx, v)
		if err != nil {
			return nil, err
		}
		result = append(result, *user)
	}
	return result, nil
}
