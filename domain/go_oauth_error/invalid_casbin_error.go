package gooautherror

type InvalidCasbinAccess struct{}

func (InvalidCasbinAccess) Error() string {
	return "INVALID_CASBIN_ACCESS"
}
