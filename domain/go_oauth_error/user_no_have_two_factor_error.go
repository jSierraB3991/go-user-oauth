package gooautherror

type UserNoHaveTwoFactorError struct{}

func (UserNoHaveTwoFactorError) Error() string {
	return "USER_NO_HAVE_TWO_FACTOR_ERROR"
}
