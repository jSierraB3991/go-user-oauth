package gooauthservice

import (
	"context"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) ChangePassword(ctx context.Context, req gooauthrequest.ChangePasswordRequest) error {
	dataUser, err := s.repo.GetUserById(jsierralibs.GetUNumberForString(req.KeycloakUserId))
	if err != nil {
		return err
	}

	isVerify := s.passwordService.VerifyPassword(dataUser.Password, req.PrePassword)
	if !isVerify {
		return gooautherror.InvalidUserOrPassword{}
	}

	encryptPasword, err := s.passwordService.EncryptPassword(req.NewPassword)
	if err != nil {
		return err
	}
	dataUser.Password = encryptPasword
	return s.repo.UpdateUser(dataUser)
}
