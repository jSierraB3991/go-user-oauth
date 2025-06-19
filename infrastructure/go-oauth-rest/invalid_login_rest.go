package gooauthrest

import "time"

type InvalidLoginRest struct {
	InvalidId        uint      `json:"invalid_id"`
	EmailLoginFailed string    `json:"email_failed"`
	IpFailed         string    `json:"ip_failed"`
	Motive           string    `json:"motive"`
	Date             time.Time `json:"date"`
	UserAgent        string    `json:"user_agent"`
	IsForTwoFactor   bool      `json:"is_two_factor"`
}

type InvalidLoginRestPagg struct {
	Limit      int                `json:"limit"`
	Page       int                `json:"page"`
	TotalRows  int64              `json:"total_rows"`
	TotalPages int                `json:"total_pages"`
	Sort       string             `json:"-"`
	Data       []InvalidLoginRest `json:"data"`
}
