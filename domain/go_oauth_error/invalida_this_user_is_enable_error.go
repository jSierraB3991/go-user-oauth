package gooautherror

type InvalidThisUserIsEnableError struct{}

func (InvalidThisUserIsEnableError) Error() string {
	return "INVALID_THIS_USER_IS_ENABLED"
}
