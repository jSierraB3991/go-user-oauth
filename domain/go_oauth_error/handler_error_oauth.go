package gooautherror

import (
	"net/http"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func GetErrorCodeUnauthorized(message string) *int {
	return eliotlibs.RunMultipleValidationCode(message, http.StatusUnauthorized,
		InactiveTokenError{},
		InvalidTokenError{},
		SessionExpiredError{},
		NotFoundSessionError{},
		QrExpiredError{},
	)
}

func GetErrorCodeForbidden(message string) *int {
	return eliotlibs.RunMultipleValidationCode(message, http.StatusForbidden,
		InvalidCasbinAccess{},
		NotFoundSessionByRefreshTokenError{},
	)
}
