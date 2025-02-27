package gooautherror

type ThePasswordIsVoidError struct{}

func (ThePasswordIsVoidError) Error() string {
	return "THE_PASSWORD_VOID_ERROR"
}

type ThePasswordIsLessToSixWordError struct{}

func (ThePasswordIsLessToSixWordError) Error() string {
	return "THE_PASSWORD_IS_LESS_SIX_WORD_ERROR"
}
