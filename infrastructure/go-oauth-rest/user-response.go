package gooauthrest

type User struct {
	Email      string               `json:"email"`
	Name       string               `json:"name"`
	SubName    string               `json:"sub_name"`
	Enabled    bool                 `json:"enabled"`
	Password   string               `json:"password"`
	Role       string               `json:"role"`
	Attributes *map[string][]string `json:"attributtes"`
}
