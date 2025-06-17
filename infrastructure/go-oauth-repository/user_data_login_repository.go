package gooauthrepository

import (
	"context"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
)

func (repo *Repository) SaveDataLogin(ctx context.Context, dataLogin gooauthmodel.UserDataLogin) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Save(&dataLogin).Error
}
