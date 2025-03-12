package gooauthrepository

import (
	"log"
	"time"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (repo *Repository) RunMigrations() error {
	err := repo.Migrate00()
	if err != nil {
		return err
	}
	err = repo.RunMigrate("01", repo.MigrateO1)
	if err != nil {
		return err
	}

	err = repo.RunMigrate("02", repo.Migrate02)
	if err != nil {
		return err
	}
	return nil
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

func (repo *Repository) Migrate02() error {
	return repo.CapitalizeNameInDatabase(1, 10)
}

func (repo *Repository) CapitalizeNameInDatabase(page, limit int) error {
	if limit < page {
		return nil
	}
	pagination := jsierralibs.Paggination{Limit: 10, Page: page}

	userDataDb, err := repo.GetUsersPage(&pagination)
	if err != nil {
		return err
	}
	for _, v := range userDataDb {
		newName := jsierralibs.CapitalizeName(v.Name)
		newSubName := jsierralibs.CapitalizeName(v.SubName)

		v.Name = newName
		v.SubName = newSubName
		err = repo.UpdateUser(&v)
		if err != nil {
			return err
		}
	}
	return repo.CapitalizeNameInDatabase(page+1, pagination.TotalPages)
}
