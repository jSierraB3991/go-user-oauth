package gooauthmodel

import (
	"time"

	"gorm.io/gorm"
)

type UserMigration struct {
	gorm.Model
	MigrationId      uint      `gorm:"column:id;not null"`
	MigrationVersion string    `gorm:"column:migrate_version;not null"`
	DateCreate       time.Time `gorm:"column:date_create;not null"`
}
