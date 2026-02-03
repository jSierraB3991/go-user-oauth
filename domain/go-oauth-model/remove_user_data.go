package gooauthmodel

import "gorm.io/gorm"

type UserDataRemove struct {
	gorm.Model
	UserDataRemoveId uint   `gorm:"column:id"`
	DataUserPpal     string `gorm:"column:data_user_ppal;not null"`
	DataUserAttr     string `gorm:"column:data_user_attr;not null"`
}
