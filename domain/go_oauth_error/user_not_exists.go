package gooautherror

type ThisUserNotExistsError struct{}

func (ThisUserNotExistsError) Error() string {
	return "THIS_USER_NOT_EXISTS_ERROR"
}
