package gooauthrequest

type ValidateOauthCodeRequest struct {
	Username string `json:"-"`
	Code     string `json:"code"`
}
