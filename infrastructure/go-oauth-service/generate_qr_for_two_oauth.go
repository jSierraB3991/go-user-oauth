package gooauthservice

import (
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) GenerateQrForDobleOuath(userName string) (*gooauthrest.QrTwoOauthRest, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MiAplicacion",
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

	return &gooauthrest.QrTwoOauthRest{
		Secret: key.Secret(),
		Url:    key.URL(),
	}, nil
}
