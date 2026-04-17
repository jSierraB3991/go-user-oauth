package gooauthservice

import (
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func (s *GoOauthService) LoginUser(ctx context.Context, req gooauthrequest.GoLoginRequest) (*gooauthrest.JWT, error) {
	userName := strings.ToLower(req.UserName)

	user, err := s.repo.GetUserByEmail(ctx, userName)
	if err != nil {
		if s.saveLoginHistory {
			s.saveInvalidDataLogin(ctx, req.Ip, req.UserAgent, userName, "Usuario no encontrado", false)
		}
		return nil, err
	}

	if !user.Enabled {
		if s.saveLoginHistory {
			s.saveInvalidDataLogin(ctx, req.Ip, req.UserAgent, userName, "Usuario invalido, no esta habilitado", false)
		}
		return nil, gooautherror.UserNotEnabledError{}
	}

	isVerify := s.passwordService.VerifyPassword(user.Password, req.Password)
	if !isVerify {
		if s.saveLoginHistory {
			s.saveInvalidDataLogin(ctx, req.Ip, req.UserAgent, userName, "La contraseña introducida, es erronea", false)
		}
		return nil, gooautherror.InvalidUserOrPassword{}
	}

	refreshToken := generateRefreshToken()
	var tokenTwoFactor, tokenTwoFactorEncrypt *string
	if user.IsActiveTwoFactorOauth {
		uuidGenerate := uuid.NewString()
		tokenTwoFactor = &uuidGenerate

		tokenTwoFactorEncryptL := s.hashToken(uuidGenerate)
		tokenTwoFactorEncrypt = &tokenTwoFactorEncryptL
	}

	sessionId, err := s.saveDataLogin(ctx, req.Ip, req.UserAgent, s.hashToken(refreshToken), tokenTwoFactorEncrypt, user.UserId, user.IsActiveTwoFactorOauth)
	if err != nil {
		log.Printf("ERROR: SAING DATA LOGIN %v", err)
		return nil, err
	}

	tokenString, exp, err := s.GetJwtToken(ctx, user.UserId, user.GoUserRoleId, user.Email, user.GoUserRole.RoleName, sessionId, req.IsRemenber)
	if err != nil {
		if s.saveLoginHistory {
			s.saveInvalidDataLogin(ctx, req.Ip, req.UserAgent, userName, "Error al generar el token", false)
		}
		return nil, err
	}

	if user.IsActiveTwoFactorOauth {
		return &gooauthrest.JWT{
			IsTwoFactor:    true,
			TokenTwoFactor: tokenTwoFactor,
		}, nil
	}

	return &gooauthrest.JWT{
		AccessToken:  tokenString,
		RefreshToken: refreshToken,
		ExpiredIn:    exp,
		Role:         user.GoUserRole.RoleName,
		IsTwoFactor:  false,
	}, nil
}
