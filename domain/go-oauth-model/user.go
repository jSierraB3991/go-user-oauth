package gooauthmodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   uint   `gorm:"column:id"`
	Email    string `gorm:"column:email;not null"`
	Name     string `gorm:"column:name;not null"`
	SubName  string `gorm:"column:sub_name"`
	Enabled  bool   `gorm:"column:enabled;not null"`
	Password string `gorm:"column:password;not null"`

	RoleId uint `gorm:"column:role_id;not null"`
	Role   Role

	CodeValidateEmail    *string `gorm:"column:code_validate_email"`
	CodeRemenberPassword *string `gorm:"column:code_remenber_password"`
}

type UserAttributtes struct {
	gorm.Model
	UserAttributteId uint `gorm:"column:id"`
	UserId           uint `gorm:"column:user_id;not null"`
	User             User
	NameAttributte   string `gorm:"column:name_attributte;not null"`
	ValueAttributtes string `gorm:"column:vale_attributte;not null"`
}

type UserPath struct {
	gorm.Model
	UserPathId uint `gorm:"column:id"`
	UserId     uint `gorm:"column:user_id;not null"`
	User       User
	PathBackId uint `gorm:"column:path_back_id;not null"`
	PathBack   PathBack
}
