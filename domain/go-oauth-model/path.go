package gooauthmodel

import "gorm.io/gorm"

type GoUserPathBack struct {
	gorm.Model
	PathBackId     uint   `gorm:"column:id"`
	PathRoute      string `gorm:"column:path_route;not null;uniqueIndex:idx_path_operation"`
	OperationRoute string `gorm:"column:operation_route;not null;uniqueIndex:idx_path_operation"`
}

type GoUserUserPath struct {
	gorm.Model
	UserPathId       uint `gorm:"column:id"`
	GoUserUserId     uint `gorm:"column:user_id;not null"`
	GoUserUser       GoUserUser
	GoUserPathBackId uint `gorm:"column:path_back_id;not null"`
	GoUserPathBack   GoUserPathBack
}
