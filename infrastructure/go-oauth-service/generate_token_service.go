package gooauthservice

import (
	"time"

	"github.com/golang-jwt/jwt"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) GeneratetokenToValidate(userId, keyToGenerateToken string, limitInHours time.Duration) (*string, error) {

	dataUser, err := s.repo.GetUserById(jsierralibs.GetUNumberForString(userId))
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dataUser.UserId,
		"exp":     time.Now().UTC().Add(limitInHours * time.Hour).Unix(), // Expira en 24 horas
	})

	var secretKey = []byte(keyToGenerateToken)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	err = s.repo.UpdateTokenMailValidatePassword(dataUser.UserId, tokenString)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
