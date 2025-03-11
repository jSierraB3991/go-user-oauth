package gooautherror

type QrExpiredError struct{}

func (QrExpiredError) Error() string {
	return "QR_EXPIRED_ERROR"
}
