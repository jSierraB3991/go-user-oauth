package gooauthservice

import (
	"context"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
)

func (s *GoOauthService) RemenberPassword(ctx context.Context, token, newPassword, codeTwoFactor string) error {
	userData, err := s.repo.GetUserByToken(ctx, token)
	if err != nil {
		return err
	}

	if userData.IsActiveTwoFactorOauth {
		isValidCode := validateInterCode(codeTwoFactor, userData.KeyOathApp, s.aesKeyForEncrypt)
		if !isValidCode {
			return gooautherror.InvalidCodeTwoFactorOauthError{}
		}
	}

	encryptPasword, err := s.passwordService.EncryptPassword(newPassword)
	if err != nil {
		return err
	}
	userData.Password = encryptPasword
	userData.TokenChangePassword = ""
	return s.repo.UpdateUser(ctx, userData)
}

func (s *GoOauthService) IsActiveTwoFactorOauth(ctx context.Context, token string) (bool, error) {
	userData, err := s.repo.GetUserByToken(ctx, token)
	if err != nil {
		return false, err
	}
	return userData.IsActiveTwoFactorOauth, nil
}
