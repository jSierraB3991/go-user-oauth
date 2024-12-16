package gooauthmodel

import "gorm.io/gorm"

type GoUserRole struct {
	gorm.Model
	RoleId   uint   `gorm:"column:id"`
	RoleName string `gorm:"column:role_name;not null"`
}

type GoUserRolePath struct {
	gorm.Model
	RolePathId uint `gorm:"column:id"`

	GoUserRoleId uint `gorm:"column:role_id;not null"`
	GoUserRole   GoUserRole

	GoUserPathBackId uint `gorm:"column:path_back_id;not null"`
	GoUserPathBack   GoUserPathBack
}

type GoUserPathBack struct {
	gorm.Model
	PathBackId     uint   `gorm:"column:id"`
	PathRoute      string `gorm:"column:path_route;not null"`
	OperationRoute string `gorm:"column:operation_route;not null"`
}
