package gooauthservice

import (
	"context"
	"strconv"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
)

func (s *GoOauthService) ValidateMailByUserId(ctx context.Context, userId string) error {
	userInt, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	user, err := s.repo.GetUserById(uint(userInt))
	if err != nil {
		return err
	}
	if user.Enabled {
		return gooautherror.InvalidThisUserIsEnableError{}
	}

	user.Enabled = true
	return s.repo.EnableUser(user.UserId)
}
