package gooautherror

type InvalidTokenError struct{}

func (InvalidTokenError) Error() string {
	return "INVALID_TOKEN_ERROR"
}
