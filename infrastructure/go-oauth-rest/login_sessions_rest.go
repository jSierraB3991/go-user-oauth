package gooauthrest

import "time"

type LoginSessionRest struct {
	Id                  uint      `json:"id"`
	UserAgent           string    `json:"user_agent"`
	Ip                  string    `json:"ip"`
	IpResponse          string    `json:"ip_response"`
	IsLoginWithPassword bool      `json:"is_login_with_password"`
	IsAvailable         bool      `json:"is_available"`
	Fecha               time.Time `json:"fecha"`
}
