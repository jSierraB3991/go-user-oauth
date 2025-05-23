package gooauthservice

import (
	"context"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) LoginUser(ctx context.Context, userName, password string) (*gooauthrest.JWT, error) {

	user, err := s.repo.GetUserByEmail(ctx, userName)
	if err != nil {
		return nil, err
	}

	if !user.Enabled {
		return nil, gooautherror.UserNotEnabledError{}
	}

	isVerify := s.passwordService.VerifyPassword(user.Password, password)
	if !isVerify {
		return nil, gooautherror.InvalidUserOrPassword{}
	}

	exp := s.GetExp()

	tokenString, err := s.GetJwtToken(ctx, exp, user.UserId, user.GoUserRoleId, user.Email, user.GoUserRole.RoleName)
	if err != nil {
		return nil, err
	}

	if user.IsActiveTwoFactorOauth {
		return &gooauthrest.JWT{
			IsTwoFactor: true,
		}, nil
	}

	return &gooauthrest.JWT{
		AccessToken:  tokenString,
		RefreshToken: tokenString,
		ExpiredIn:    exp,
		Role:         user.GoUserRole.RoleName,
		IsTwoFactor:  false,
	}, nil
}
