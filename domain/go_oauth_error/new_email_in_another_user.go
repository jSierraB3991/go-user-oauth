package gooautherror

type NewEmailInAntherUserError struct{}

func (NewEmailInAntherUserError) Error() string {
	return "NEW_EMAIL_IN_ANOTHER_USER"
}
