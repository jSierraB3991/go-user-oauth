package gooauthservice

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
)

func (s *GoOauthService) GenerateValidateMail(mailSend string) (string, error) {
	userData, err := s.repo.GetUserByEmail(mailSend)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userData.UserId,
		"exp":     time.Now().UTC().Add(24 * time.Hour).Unix(), // Expira en 24 horas
	})

	var secretKey = []byte(s.secretForJwt)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	err = s.repo.UpdateLinkMailValidateMail(userData.UserId, tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *GoOauthService) ValidateMailByUserId(ctx context.Context, userId string) error {
	userInt, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	user, err := s.repo.GetUserById(uint(userInt))
	if err != nil {
		return err
	}
	if user.Enabled {
		return gooautherror.InvalidThisUserIsEnableError{}
	}

	user.Enabled = true
	user.LinkToValidateMail = ""
	return s.repo.EnableUser(user.UserId)
}
