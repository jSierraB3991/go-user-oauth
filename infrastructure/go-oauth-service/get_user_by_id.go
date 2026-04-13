package gooauthservice

import (
	"context"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) GetUserByUserId(ctx context.Context, userId string) (*gooauthrest.User, error) {
	userInt := eliotlibs.GetUNumberForString(userId)
	if userInt == 0 {
		return nil, gooautherror.InvalidUserIdError{UserId: userId}
	}
	user, err := s.repo.GetUserById(ctx, userInt)
	if err != nil {
		return nil, err
	}

	attributtes, err := s.repo.GetAttributtesByUserId(ctx, userInt)
	if err != nil {
		return nil, err
	}

	return gooauthmapper.GetUserRestAnAttributtes(user, attributtes), nil
}

func (s *GoOauthService) GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error) {
	usersId, err := gooauthlibs.GetUintsFromStrings(keycloakUsersId)
	if err != nil {
		return nil, err
	}
	users, err := s.repo.GetUsersByIds(ctx, usersId)
	if err != nil {
		return nil, err
	}

	attributtes, err := s.repo.GetAttributtesByUserIds(ctx, usersId)
	if err != nil {
		return nil, err
	}
	return gooauthmapper.GetUsersRestAnAttributtes(users, attributtes), nil
}
