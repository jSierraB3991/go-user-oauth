package gooauthservice

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func generateRefreshToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func (s *GoOauthService) hashToken(token string) string {
	h := hmac.New(sha256.New, []byte(s.secretForJwt))
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *GoOauthService) RefreshToken(ctx context.Context, refreshToken string) (*gooauthrest.JWT, error) {
	refreshTokenE := s.hashToken(refreshToken)

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
	newRefreshTokenEncrypt := s.hashToken(newRefreshToken)

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
