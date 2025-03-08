package gooauthservice

import (
	"strings"

	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) GenerateQrForDobleOuath(userName string) (*gooauthrest.QrTwoOauthRest, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.appName,
		AccountName: userName,
	})
	if err != nil {
		return nil, err
	}

	secretEncrypt, err := jsierralibs.Encrypt(key.Secret(), s.aesKeyForEncrypt)
	if err != nil {
		return nil, err
	}
	err = s.repo.SaveSecretToUser(userName, secretEncrypt)
	if err != nil {
		return nil, err
	}

	result := key.URL()

	if strings.Trim(s.urlImagenApp, " ") != "" {
		result += "&image=" + s.urlImagenApp
	}

	return &gooauthrest.QrTwoOauthRest{
		Url: result,
	}, nil
}

func (s *GoOauthService) IsActiveTwoFactor(user string) (bool, error) {
	userData, err := s.repo.GetUserByEmail(user)
	if err != nil {
		return false, err
	}
	return userData.IsActiveTwoFactorOauth, nil
}
