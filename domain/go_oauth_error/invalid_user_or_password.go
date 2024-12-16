package gooautherror

type InvalidUserOrPassword struct{}

func (InvalidUserOrPassword) Error() string {
	return "INVALID_USER_OR_PASSWORD"
}
