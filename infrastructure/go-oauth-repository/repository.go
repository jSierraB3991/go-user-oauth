package gooauthrepository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func InitiateRepo(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
