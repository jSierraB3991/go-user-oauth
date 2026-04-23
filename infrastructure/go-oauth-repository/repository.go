package gooauthrepository

import (
	"context"
	"fmt"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	"gorm.io/gorm"
)

type Repository struct {
	db                  *gorm.DB
	schemaForMigrations string
}

func InitiateRepo(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (repo *Repository) WithTenant(ctx context.Context) (*gorm.DB, error) {
	tenant, err := eliotlibs.WithTenant(ctx)
	if err != nil {
		return repo.db, nil
	}
	if repo.db == nil {
		return nil, gooautherror.NotDatabaseConfigurateError{}
	}

	tx := repo.db.Session(&gorm.Session{
		NewDB:       true,
		PrepareStmt: false,
	})

	var currentSearchPath string
	if err := tx.Raw("SHOW search_path").Scan(&currentSearchPath).Error; err != nil {
		return nil, err
	}

	if "\""+currentSearchPath+"\"" == *tenant || currentSearchPath == *tenant {
		return tx, nil
	}

	fmt.Println("Switching schema to:", *tenant)
	if err := tx.Exec(fmt.Sprintf(`SET search_path TO %s`, *tenant)).Error; err != nil {
		return nil, err
	}

	return tx, nil
}
