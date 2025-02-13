package gooauthmodel

import "gorm.io/gorm"

type GoUserRole struct {
	gorm.Model
	RoleId   uint   `gorm:"column:id"`
	RoleName string `gorm:"column:role_name;not null;unique"`
}

type GoUserRolePath struct {
	gorm.Model
	RolePathId uint `gorm:"column:id"`

	GoUserRoleId uint `gorm:"column:role_id;not null"`
	GoUserRole   GoUserRole

	GoUserPathBackId uint `gorm:"column:path_back_id;not null"`
	GoUserPathBack   GoUserPathBack
}
