package gooauthservice

import (
	"context"
	"strings"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) ChangeEmailByAdmin(ctx context.Context, kUserId, newEmail string) error {
	userId := eliotlibs.GetUNumberForString(kUserId)
	user, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	if user.UserId != userId {
		return gooautherror.ThisUserNotExistsError{}
	}

	user.Email = strings.ToLower(newEmail)
	return s.repo.UpdateUser(ctx, user)
}
