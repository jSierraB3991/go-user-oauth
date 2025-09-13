package gooauthrepository

import (
	"context"
	"fmt"
	"log"
	"time"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	"gorm.io/gorm"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (repo *Repository) RunMigrations(schemas []string) error {

	for _, schema := range schemas {
		// Asegúrate de que el schema existe antes de migrar
		if err := ensureSchemaExists(repo.db, schema); err != nil {
			return fmt.Errorf("schema '%s' creation failed: %w", schema, err)
		}

		dbTenant, err := repo.db.Session(&gorm.Session{NewDB: true}).DB()
		if err != nil {
			log.Fatalf("could not create session for %s: %v", schema, err)
		}

		// establecer el search_path de forma segura
		_, err = dbTenant.Exec(`SET search_path TO ` + eliotlibs.QuoteIdentifier(schema)) // o con una query preparada
		if err != nil {
			log.Fatalf("could not set search_path for %s: %v", schema, err)

		}

		repo.schemaForMigrations = schema

		err = repo.Migrate00()
		if err != nil {
			return err
		}
		eliotlibs.RunMigrations(repo,
			repo.MigrateO1,
			repo.Migrate02)
		log.Printf("✅ Migrated schema: %s", schema)
	}
	return nil
}

func ensureSchemaExists(db *gorm.DB, schema string) error {
	return db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", eliotlibs.QuoteIdentifier(schema))).Error
}

func (repo *Repository) ValidateMigrate(version string) (bool, error) {
	var result gooauthmodel.GoUserUserMigration
	err := repo.db.Where("migrate_version = ?", version).Find(&result).Error
	if err != nil {
		return true, err
	}

	if result.MigrationVersion == version {
		return true, nil
	}
	return false, nil
}

func (repo *Repository) SaveMigration(version string) error {
	return repo.db.Save(&gooauthmodel.GoUserUserMigration{DateCreate: time.Now(), MigrationVersion: version}).Error
}

func (repo *Repository) Migrate00() error {
	return repo.db.AutoMigrate(
		&gooauthmodel.GoUserDataLogin{},
		&gooauthmodel.GoUserInvalidGoAuth{},
		&gooauthmodel.GoUserUserMigration{},
		&gooauthmodel.GoUserPathBack{},
		&gooauthmodel.GoUserRole{},
		&gooauthmodel.GoUserRolePath{},
		&gooauthmodel.GoUserUser{},
		&gooauthmodel.GoUserUserAttributtes{},
		&gooauthmodel.GoUserUserPath{},
	)
}

func (repo *Repository) Migrate02() error {
	ctx := context.WithValue(context.Background(), eliotlibs.ContextTenantKey, repo.schemaForMigrations)
	return repo.capitalizeNameInDatabase(ctx, 1, 10)
}

func (repo *Repository) capitalizeNameInDatabase(ctx context.Context, page, limit int) error {
	if limit < page {
		return nil
	}
	pagination := eliotlibs.Paggination{Limit: 10, Page: page}

	userDataDb, err := repo.GetUsersPage(ctx, &pagination)
	if err != nil {
		return err
	}
	for _, v := range userDataDb {
		newName := eliotlibs.CapitalizeName(v.Name)
		newSubName := eliotlibs.CapitalizeName(v.SubName)

		v.Name = newName
		v.SubName = newSubName
		err = repo.UpdateUser(ctx, &v)
		if err != nil {
			return err
		}
	}
	return repo.capitalizeNameInDatabase(ctx, page+1, pagination.TotalPages)
}
