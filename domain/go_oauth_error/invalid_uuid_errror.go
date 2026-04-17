package gooautherror

type InvalidUuuidTokenError struct{}

func (InvalidUuuidTokenError) Error() string {
	return "INVALID_UUID_TOKEN_ERROR"
}

type TokenUuidExpiredError struct{}

func (TokenUuidExpiredError) Error() string {
	return "TOKEN_UUID_EXPIRED_ERROR"
}
