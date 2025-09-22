package gooauthrest

type User struct {
	Id                     string               `json:"id"`
	Email                  string               `json:"email"`
	Name                   string               `json:"name"`
	SubName                string               `json:"sub_name"`
	Enabled                bool                 `json:"enabled"`
	Password               string               `json:"password"`
	Role                   string               `json:"role"`
	IsActiveTwoFactorOauth bool                 `json:"is_active_two_factor_oauth"`
	Attributes             *map[string][]string `json:"attributtes"`
}
