package gooauthservice

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) ValidateCodeOtp(ctx context.Context, req gooauthrequest.ValidateOauthCodeRequest) (bool, error) {
	userName := strings.ToLower(req.Username)
	user, err := s.repo.GetUserByEmail(ctx, userName)
	if err != nil {
		return false, err
	}

	code := user.KeyOathApp

	secretData, err := eliotlibs.Decrypt(code, s.aesKeyForEncrypt)
	if err != nil {
		return false, err
	}

	parts := strings.Split(secretData, "|")
	if len(parts) != 2 {
		return false, errors.New("invalid secret format")
	}

	codeDecrypeted := parts[0]
	expirationTime, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return false, err
	}

	if time.Now().Unix() > expirationTime {
		return false, gooautherror.QrExpiredError{}
	}

	isValidCode := totp.Validate(req.Code, codeDecrypeted)

	if isValidCode && !user.IsActiveTwoFactorOauth {
		err = s.repo.ActiveTwoFactorOauth(ctx, userName)
		if err != nil {
			return false, err
		}
	}

	return isValidCode, nil
}

func (s *GoOauthService) ValidateCodeTwoFactor(ctx context.Context, req gooauthrequest.ValidateOauthCodeRequest) (bool, error) {
	userName := strings.ToLower(req.Username)
	user, err := s.repo.GetUserByEmail(ctx, userName)
	if err != nil {
		return false, err
	}

	return validateInterCode(req.Code, user.KeyOathApp, s.aesKeyForEncrypt), nil
}

func validateInterCode(reqCode, userKeyOauth, aesKey string) bool {

	secretData, err := eliotlibs.Decrypt(userKeyOauth, aesKey)
	if err != nil {
		return false
	}

	parts := strings.Split(secretData, "|")
	if len(parts) != 2 {
		return false
	}

	return totp.Validate(reqCode, parts[0])
}
