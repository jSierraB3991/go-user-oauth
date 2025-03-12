package gooauthservice

import (
	"strings"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
)

func (s *PasswordService) ValidatePassword(password string) error {
	err := validateLength(password)
	if err != nil {
		return err
	}
	return nil
}

func validateLength(password string) error {
	if strings.TrimSpace(password) == "" {
		return gooautherror.ThePasswordIsVoidError{}
	}

	if len(strings.TrimSpace(password)) < 6 {
		return gooautherror.ThePasswordIsLessToSixWordError{}
	}
	return nil
}
