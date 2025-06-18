package gooauthrepository

import (
	"context"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
)

func (repo *Repository) SaveDataLogin(ctx context.Context, dataLogin gooauthmodel.GoUserDataLogin) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Save(&dataLogin).Error
}

func (repo *Repository) SaveInvalidLogin(ctx context.Context, invalidDataLogin gooauthmodel.GoUserInvalidGoAuth) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Save(&invalidDataLogin).Error
}
