package gooauthservice

import (
	"context"
	"log"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) RemoveSessionByRefreshToken(ctx context.Context, email, refreshToken string) error {
	refreshTokenE, err := eliotlibs.Encrypt(refreshToken, s.aesKeyForEncrypt)
	if err != nil {
		return err
	}
	sessions, err := s.repo.GetSessionsByEmailRefreshTokenE(ctx, email, refreshTokenE)
	if err != nil {
		return err
	}

	if len(sessions) <= 0 {
		return nil
	}

	for _, v := range sessions {
		refreshTokenDecrypr, err := eliotlibs.Decrypt(v.RefreshToken, s.aesKeyForEncrypt)
		if err != nil {
			return err
		}

		if refreshTokenDecrypr == refreshToken {
			err := s.repo.RemoveSessionById(ctx, v.UserDataLoginId)
			if err != nil {
				log.Println(err)
			}
			break
		}
	}
	return nil
}

func (s *GoOauthService) RemoveSessionById(ctx context.Context, sessionId uint) error {
	return s.repo.RemoveSessionById(ctx, sessionId)
}

func (s *GoOauthService) RemoveOldSessions(ctx context.Context) {
	err := s.repo.RemoveSessionsPreDate(ctx)
	if err != nil {
		log.Printf("ERROR CLEANING SESSIONS: %s", err.Error())
	}
}
