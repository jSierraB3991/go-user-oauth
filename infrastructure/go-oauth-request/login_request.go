package gooauthrequest

type GoLoginRequest struct {
	UserName   string
	Password   string
	Ip         string
	UserAgent  string
	IsRemenber bool
}

type GoLoginRequestTwoFactor struct {
	UserName      string
	CodeTwoFactor string
	Ip            string
	UserAgent     string
	IsRemenber    bool
}
