package gooautherror

type InvalidRole struct{}

func (InvalidRole) Error() string {
	return "INVALID_ROLE"
}
