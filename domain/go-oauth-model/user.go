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
}

type GoUserUserAttributtes struct {
	gorm.Model
	UserAttributteId uint `gorm:"column:id"`
	GoUserUserId     uint `gorm:"column:user_id;not null"`
	GoUserUser       GoUserUser
	NameAttributte   string `gorm:"column:name_attributte;not null"`
	ValueAttributtes string `gorm:"column:vale_attributte;not null"`
}

type GoUserUserPath struct {
	gorm.Model
	UserPathId       uint `gorm:"column:id"`
	GoUserUserId     uint `gorm:"column:user_id;not null"`
	GoUserUser       GoUserUser
	GoUserPathBackId uint `gorm:"column:path_back_id;not null"`
	GoUserPathBack   GoUserPathBack
}
