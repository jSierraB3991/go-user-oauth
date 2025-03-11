package gooauthservice

import (
	"context"
	"errors"
	"strings"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) LoginWithTwoFactor(ctx context.Context, userName, codeTwoFactor string) (*gooauthrest.JWT, error) {
	user, err := s.repo.GetUserByEmail(userName)
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
	isValidCode := totp.Validate(codeTwoFactor, codeDecrypeted)
	if !isValidCode {
		return nil, gooautherror.InvalidCodeTwoFactorOauthError{}
	}

	exp := s.GetExp()
	tokenString, err := s.GetJwtToken(exp, user.UserId, user.Email)
	if err != nil {
		return nil, err
	}
	return &gooauthrest.JWT{
		AccessToken:  tokenString,
		RefreshToken: tokenString,
		ExpiredIn:    exp,
		Role:         user.GoUserRole.RoleName,
	}, nil
}
