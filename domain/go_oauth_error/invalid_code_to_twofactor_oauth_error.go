package gooautherror

type InvalidCodeTwoFactorOauthError struct{}

func (InvalidCodeTwoFactorOauthError) Error() string {
	return "INVALID_CODE_TWO_FACTOR_OAUTH_EROR"
}
