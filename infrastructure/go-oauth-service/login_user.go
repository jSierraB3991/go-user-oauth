package gooauthservice

import (
	"context"
	"log"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) LoginUser(ctx context.Context, req gooauthrequest.GoLoginRequest) (*gooauthrest.JWT, error) {

	user, err := s.repo.GetUserByEmail(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	if !user.Enabled {
		return nil, gooautherror.UserNotEnabledError{}
	}

	isVerify := s.passwordService.VerifyPassword(user.Password, req.Password)
	if !isVerify {
		return nil, gooautherror.InvalidUserOrPassword{}
	}

	tokenString, exp, err := s.GetJwtToken(ctx, user.UserId, user.GoUserRoleId, user.Email, user.GoUserRole.RoleName, req.IsRemenber)
	if err != nil {
		return nil, err
	}

	if user.IsActiveTwoFactorOauth {
		return &gooauthrest.JWT{
			IsTwoFactor: true,
		}, nil
	}

	err = s.saveDataLogin(ctx, req.Ip, req.UserAgent, tokenString, user.UserId, true)
	if err != nil {
		log.Printf("ERROR: SAING DATA LOGIN %v", err)
	}

	return &gooauthrest.JWT{
		AccessToken:  tokenString,
		RefreshToken: tokenString,
		ExpiredIn:    exp,
		Role:         user.GoUserRole.RoleName,
		IsTwoFactor:  false,
	}, nil
}
