package gooauthservice

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func generateRefreshToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func (s *GoOauthService) RefreshToken(ctx context.Context, refreshToken string) (*gooauthrest.JWT, error) {
	refreshTokenE, err := eliotlibs.Encrypt(refreshToken, s.aesKeyForEncrypt)
	if err != nil {
		return nil, err
	}

	session, err := s.repo.GetSessionsByRefreshToken(ctx, refreshTokenE)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, gooautherror.NotFoundSessionByRefreshTokenError{}
	}

	if session.ExpiresAt.Before(time.Now()) {
		err = s.repo.RemoveSessionById(ctx, session.UserDataLoginId)
		if err != nil {
			return nil, err
		}
		return nil, gooautherror.SessionExpiredError{}
	}

	newRefreshToken := generateRefreshToken()
	newRefreshTokenEncrypt, err := eliotlibs.Encrypt(newRefreshToken, s.aesKeyForEncrypt)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateRefreshToken(ctx, session.UserDataLoginId, newRefreshTokenEncrypt)
	if err != nil {
		return nil, err
	}

	tokenString, exp, err := s.GetJwtToken(ctx, session.GoUserUserId,
		session.GoUserUser.GoUserRoleId, session.GoUserUser.Email,
		session.GoUserUser.GoUserRole.RoleName, session.UserDataLoginId, false)
	if err != nil {
		return nil, err
	}

	return &gooauthrest.JWT{
		AccessToken:  tokenString,
		RefreshToken: newRefreshToken,
		ExpiredIn:    exp,
		Role:         session.GoUserUser.GoUserRole.RoleName,
		IsTwoFactor:  false,
	}, nil
}
