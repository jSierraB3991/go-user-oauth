package gooauthrepository

import (
	"log"
	"time"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
)

func (repo *Repository) RunMigrations() error {
	err := repo.Migrate00()
	if err != nil {
		return err
	}
	return repo.RunMigrate("01", repo.MigrateO1)
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

func (repo *Repository) SaveVersion(version string) error {
	return repo.db.Save(&gooauthmodel.GoUserUserMigration{DateCreate: time.Now(), MigrationVersion: version}).Error
}

func (repo *Repository) RunMigrate(version string, migration func() error) error {

	exist, err := repo.ValidateMigrate(version)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	err = migration()
	if err != nil {
		return err
	}
	log.Printf("SAVING MIGRATION %s", version)
	return repo.SaveVersion(version)
}

func (repo *Repository) Migrate00() error {
	return repo.db.AutoMigrate(
		&gooauthmodel.GoUserUserMigration{},
		&gooauthmodel.GoUserPathBack{},
		&gooauthmodel.GoUserRole{},
		&gooauthmodel.GoUserRolePath{},
		&gooauthmodel.GoUserUser{},
		&gooauthmodel.GoUserUserAttributtes{},
		&gooauthmodel.GoUserUserPath{},
	)
}
