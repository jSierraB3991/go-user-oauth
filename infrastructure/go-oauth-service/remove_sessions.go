package gooauthservice

import (
	"context"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) RemoveSessionByToken(ctx context.Context, email, tokenString string) error {
	sessions, err := s.repo.GetSessionsByToken(ctx, email, tokenString)
	if err != nil {
		return err
	}

	if len(sessions) <= 0 {
		return nil
	}

	for _, v := range sessions {
		tokenDecrypr, err := eliotlibs.Decrypt(v.Token, s.aesKeyForEncrypt)
		if err != nil {
			return err
		}

		if tokenDecrypr == tokenString {
			s.repo.RemoveSessionById(ctx, v.UserDataLoginId)
			break
		}
	}
	return nil
}

func (s *GoOauthService) RemoveSessionById(ctx context.Context, sessionId uint) error {
	return s.repo.RemoveSessionById(ctx, sessionId)
}
