package gooautherror

type NotFoundSessionByRefreshTokenError struct{}

func (NotFoundSessionByRefreshTokenError) Error() string {
	return "NOT_FOUND_SESSION_BY_REFRESH_TOKEN_ERROR"
}

type SessionExpiredError struct{}

func (SessionExpiredError) Error() string {
	return "SESSION_EXPIRED_ERROR"
}

type NotFoundSessionError struct{}

func (NotFoundSessionError) Error() string {
	return "NOT_FOUND_SESSION_ERROR"
}
