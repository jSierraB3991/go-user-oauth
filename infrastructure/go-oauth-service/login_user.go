package gooauthservice

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) LoginUser(ctx context.Context, userName, password string) (*gooauthrest.JWT, error) {
	user, err := s.repo.GetUserByEmail(userName)
	if err != nil {
		return nil, err
	}

	if !user.Enabled {
		return nil, gooautherror.UserNotEnabledError{}
	}

	isVerify := s.passwordService.VerifyPassword(user.Password, password)
	if !isVerify {
		return nil, gooautherror.InvalidUserOrPassword{}
	}

	pathsAllow, err := s.repo.GetPathAllowByUser(user.UserId)
	if err != nil {
		return nil, err
	}

	exp := s.GetExp()

	tokenString, err := s.GetJwtToken(exp, user.UserId, pathsAllow, user.Email)
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

func (s *GoOauthService) GetJwtToken(exp int, userId uint, pathsAllow []string, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":     userId,
		"exp":         exp,
		"paths_allow": pathsAllow,
		"iat":         time.Now().Unix(),
		"email":       email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretForJwt))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *GoOauthService) GetExp() int {
	return int(time.Now().Add(24 * time.Hour).Unix())
}
