package gooauthrepository

import (
	"context"
	"fmt"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func InitiateRepo(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (p *Repository) WithTenant(ctx context.Context) (*gorm.DB, error) {

	tenant, err := jsierralibs.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	// Creamos una nueva sesión segura
	tx := p.db.Session(&gorm.Session{
		NewDB: true,
	})
	// Ejecutar con interpolación controlada (porque no se puede parametrizar)
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", *tenant)).Error; err != nil {
		return nil, err
	}

	return tx, nil
}
