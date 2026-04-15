package gooauthservice

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
)

func (s *GoOauthService) ValidateTokenIsValidSession(ctx context.Context, tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return gooautherror.InvalidTokenError{}
	}

	claims := token.Claims.(jwt.MapClaims)
	sid, ok := claims[gooauthlibs.SESSION_ID].(uint)
	if !ok {
		return gooautherror.InvalidTokenError{}
	}

	session, err := s.repo.GetSessionById(ctx, sid)
	if err != nil {
		return err
	}

	if session == nil || !session.IsAvailable {
		return gooautherror.NotFoundSessionError{}
	}

	if session.ExpiresAt.Before(time.Now()) {
		err := s.repo.RemoveSessionById(ctx, sid)
		if err != nil {
			return err
		}
		return gooautherror.SessionExpiredError{}
	}

	return nil
}
