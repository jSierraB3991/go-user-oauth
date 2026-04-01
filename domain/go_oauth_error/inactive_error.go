package gooautherror

type InactiveTokenError struct{}

func (InactiveTokenError) Error() string {
	return "INACTIVE_TOKEN_ERROR"
}
