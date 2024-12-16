package gooauthrest

type JWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredIn    int    `json:"expired_in"`
	Role         string `json:"role"`
}
