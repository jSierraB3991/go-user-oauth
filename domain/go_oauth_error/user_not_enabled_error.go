package gooautherror

type UserNotEnabledError struct{}

func (UserNotEnabledError) Error() string {
	return "USER_NOT_ENABLED_ERROR"
}
