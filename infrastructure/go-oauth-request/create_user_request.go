package gooauthrequest

type CreateUser struct {
	Email       string `json:"email"`
	UserName    string `json:"-"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Emailverify bool   `json:"-"`
}
