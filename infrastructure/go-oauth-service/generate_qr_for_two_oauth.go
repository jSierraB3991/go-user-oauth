package gooauthservice

import (
	"context"
	"fmt"
	"strings"
	"time"

	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) GenerateQrForDobleOuath(ctx context.Context, userName, appName, imageUrl string) (*gooauthrest.QrTwoOauthRest, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      appName,
		AccountName: userName,
	})
	if err != nil {
		return nil, err
	}

	//secretEncrypt, err := jsierralibs.Encrypt(key.Secret(), s.aesKeyForEncrypt)
	expirationTime := time.Now().Add(s.timeToExpiredQrForOauth * time.Minute).Unix() // Expira en 5 minutos
	secretData := fmt.Sprintf("%s|%d", key.Secret(), expirationTime)                 // Concatenar secreto + timestamp
	secretEncrypt, err := jsierralibs.Encrypt(secretData, s.aesKeyForEncrypt)

	if err != nil {
		return nil, err
	}
	err = s.repo.SaveSecretToUser(ctx, userName, secretEncrypt)
	if err != nil {
		return nil, err
	}

	result := key.URL()

	if strings.TrimSpace(imageUrl) != "" {
		result += "&image=" + imageUrl
	}

	return &gooauthrest.QrTwoOauthRest{
		Url: result,
	}, nil
}

func (s *GoOauthService) IsActiveTwoFactor(ctx context.Context, user string) (bool, error) {
	userData, err := s.repo.GetUserByEmail(ctx, user)
	if err != nil {
		return false, err
	}
	return userData.IsActiveTwoFactorOauth, nil
}
