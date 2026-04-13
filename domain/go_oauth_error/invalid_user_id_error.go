package gooautherror

type InvalidUserIdError struct {
	UserId string
}

func (InvalidUserIdError) Error() string {
	return "INVALID_USER_ID_ERROR"
}
