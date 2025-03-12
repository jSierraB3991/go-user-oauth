package gooauthmodel

import "gorm.io/gorm"

type GoUserUser struct {
	gorm.Model
	UserId   uint   `gorm:"column:id"`
	Email    string `gorm:"column:email;not null"`
	Name     string `gorm:"column:name;not null"`
	SubName  string `gorm:"column:sub_name"`
	Enabled  bool   `gorm:"column:enabled;not null"`
	Password string `gorm:"column:password;not null"`

	GoUserRoleId uint `gorm:"column:role_id;not null"`
	GoUserRole   GoUserRole

	CodeValidateEmail    *string `gorm:"column:code_validate_email"`
	CodeRemenberPassword *string `gorm:"column:code_remenber_password"`

	KeyOathApp             string `gorm:"column:key_oath_app"`
	IsActiveTwoFactorOauth bool   `gorm:"column:is_active_two_factor;not null;default:false"`

	TokenChangePassword string `gorm:"column:token_to_change_password"`
	LinkToValidateMail  string `gorm:"column:link_to_validate_mail"`
}

type GoUserUserAttributtes struct {
	gorm.Model
	UserAttributteId uint `gorm:"column:id"`
	GoUserUserId     uint `gorm:"column:user_id;not null"`
	GoUserUser       GoUserUser
	NameAttributte   string `gorm:"column:name_attributte;not null"`
	ValueAttributtes string `gorm:"column:vale_attributte;not null"`
}
