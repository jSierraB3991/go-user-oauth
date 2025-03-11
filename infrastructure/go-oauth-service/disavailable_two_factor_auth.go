package gooauthservice

import (
	"errors"
	"strings"
	"time"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) DisAvailableTwoFactorAuth(userEmail, codeTwoFactor string) error {
	user, err := s.repo.GetUserByEmail(userEmail)
	if err != nil {
		return err
	}

	secretData, err := jsierralibs.Decrypt(user.KeyOathApp, s.aesKeyForEncrypt)
	if err != nil {
		return err
	}

	parts := strings.Split(secretData, "|")
	if len(parts) != 2 {
		return errors.New("invalid secret format")
	}

	codeDecrypeted := parts[0]
	isValidCode := totp.Validate(codeTwoFactor, codeDecrypeted)
	if !isValidCode {
		time.Sleep(5 * time.Second)
		return gooautherror.InvalidCodeTwoFactorOauthError{}
	}

	user.KeyOathApp = ""
	user.IsActiveTwoFactorOauth = false
	return s.repo.UpdateUser(user)
}
