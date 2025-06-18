package gooauthservice

import (
	"context"
	"errors"
	"log"
	"strings"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) LoginWithTwoFactor(ctx context.Context, req gooauthrequest.GoLoginRequestTwoFactor) (*gooauthrest.JWT, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	if !user.IsActiveTwoFactorOauth {
		return nil, gooautherror.UserNoHaveTwoFactorError{}
	}

	secretData, err := jsierralibs.Decrypt(user.KeyOathApp, s.aesKeyForEncrypt)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(secretData, "|")
	if len(parts) != 2 {
		return nil, errors.New("invalid secret format")
	}

	codeDecrypeted := parts[0]
	isValidCode := totp.Validate(req.CodeTwoFactor, codeDecrypeted)
	if !isValidCode {
		return nil, gooautherror.InvalidCodeTwoFactorOauthError{}
	}

	tokenString, exp, err := s.GetJwtToken(ctx, user.UserId, user.GoUserRoleId, user.Email, user.GoUserRole.RoleName, req.IsRemenber)
	if err != nil {
		return nil, err
	}

	err = s.saveDataLogin(ctx, req.Ip, req.UserAgent, tokenString, user.UserId, false)
	if err != nil {
		log.Printf("ERROR: SAING DATA LOGIN %v", err)
	}
	return &gooauthrest.JWT{
		AccessToken:  tokenString,
		RefreshToken: tokenString,
		ExpiredIn:    exp,
		Role:         user.GoUserRole.RoleName,
	}, nil
}
