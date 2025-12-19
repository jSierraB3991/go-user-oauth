package gooauthrest

import (
	"time"

	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
)

type LoginSessionRest struct {
	Id                  uint                         `json:"id"`
	UserAgent           string                       `json:"user_agent"`
	Ip                  string                       `json:"ip"`
	IpResponse          gooauthrequest.IPInfoRequest `json:"ip_response"`
	IsLoginWithPassword bool                         `json:"is_login_with_password"`
	IsAvailable         bool                         `json:"is_available"`
	Fecha               time.Time                    `json:"fecha"`
	IsThisSession       bool                         `json:"is_this_session"`
}

type LoginSessionRestPagination struct {
	Limit      int                `json:"limit"`
	Page       int                `json:"page"`
	TotalRows  int64              `json:"rows"`
	TotalPages int                `json:"pages"`
	Data       []LoginSessionRest `json:"data"`
	Sort       string             `json:"-"`
}
