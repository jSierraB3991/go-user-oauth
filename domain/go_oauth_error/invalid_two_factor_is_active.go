package gooautherror

type InvalidTwoFactorIsActive struct{}

func (InvalidTwoFactorIsActive) Error() string {
	return "INVALID_TWO_FACTOR_IS_ACTIVE"
}
