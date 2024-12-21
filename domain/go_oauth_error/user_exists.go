package gooautherror

type UserExistsError struct{}

func (UserExistsError) Error() string {
	return "THIS_EMAIL_EXISTS_IN_DATABASE"
}
