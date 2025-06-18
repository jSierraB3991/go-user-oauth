package gooauthmodel

import (
	"time"

	"gorm.io/gorm"
)

type GoUserDataLogin struct {
	gorm.Model
	UserDataLoginId     uint      `gorm:"column:id"`
	Ip                  string    `gorm:"column:ip"`
	UserAgent           string    `gorm:"column:user_agent;not null"`
	IpResponse          string    `gorm:"column:info_response"`
	IsLoginWithPassword bool      `gorm:"column:is_login_with_password;not null"`
	Token               string    `gorm:"column:token;not null"`
	GoUserUserId        uint      `gorm:"column:user_id;not null"`
	IsAvailable         bool      `gorm:"column:is_available;not null;default:true"`
	Fecha               time.Time `gorm:"column:fecha;not null"`
	GoUserUser          GoUserUser
}

type GoUserInvalidGoAuth struct {
	gorm.Model
	Email               string    `gorm:"column:email"`
	Ip                  string    `gorm:"column:ip"`
	Motive              string    `gorm:"column:ip;not null"`
	IsUtil              bool      `gorm:"column:is_util;not null"`
	Fecha               time.Time `gorm:"column:fecha;not null"`
	UserAgent           string    `gorm:"column:user_agent;not null"`
	IpResponse          string    `gorm:"column:info_response"`
	IsForTwoFactorOauth bool      `gorm:"column:is_for_two_factor_ath;not null"`
	TenantId            string    `gorm:"column:tenant_id;not null"`
}
