package gooautherror

type InactiveToken struct{}

func (InactiveToken) Error() string {
	return "INACTIVE_TOKEN"
}
