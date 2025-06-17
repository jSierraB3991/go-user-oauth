package gooauthmodel

import (
	"time"

	"gorm.io/gorm"
)

type UserDataLogin struct {
	gorm.Model
	UserDataLoginId     uint      `gorm:"column:id"`
	Ip                  string    `gorm:"column:ip"`
	UserAgent           string    `gorm:"column:user_agent;not null"`
	IpResponse          string    `gorm:"column:info_response"`
	IsLoginWithPassword bool      `gorm:"column:is_login_with_password;not null"`
	Token               string    `gorm:"column:token;not null"`
	GoUserUserId        uint      `gorm:"column:user_id;not null"`
	Fecha               time.Time `gorm:"column:fecha;not null"`
	GoUserUser          GoUserUser
}
