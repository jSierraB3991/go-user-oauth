package gooauthservice

import (
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) ValidateCodeOtp(req gooauthrequest.ValidateOauthCodeRequest) (bool, error) {
	code, err := s.repo.GetSecretOauthCode(req.Username)
	if err != nil {
		return false, err
	}

	codeDecrypeted, err := jsierralibs.Decrypt(*code, s.aesKeyForEncrypt)
	if err != nil {
		return false, err
	}

	return totp.Validate(req.Code, codeDecrypeted), nil
}
