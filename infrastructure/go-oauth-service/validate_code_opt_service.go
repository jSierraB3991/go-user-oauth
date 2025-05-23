package gooauthservice

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"github.com/pquerna/otp/totp"
)

func (s *GoOauthService) ValidateCodeOtp(ctx context.Context, req gooauthrequest.ValidateOauthCodeRequest) (bool, error) {
	code, err := s.repo.GetSecretOauthCode(ctx, req.Username)
	if err != nil {
		return false, err
	}

	secretData, err := jsierralibs.Decrypt(*code, s.aesKeyForEncrypt)
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

	if isValidCode {
		err = s.repo.ActiveTwoFactorOauth(ctx, req.Username)
		if err != nil {
			return false, err
		}
	}

	return isValidCode, nil
}
