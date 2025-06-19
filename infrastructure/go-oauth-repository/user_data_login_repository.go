package gooauthrepository

import (
	"context"
	"log"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
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

func (repo *Repository) GetDataLoginFailed(ctx context.Context, page *eliotlibs.Paggination) ([]gooauthmodel.GoUserInvalidGoAuth, error) {
	var result []gooauthmodel.GoUserInvalidGoAuth
	db, err := repo.WithTenant(ctx)
	if err != nil {
		log.Printf("ERROR GET DB %s", err.Error())
		return nil, err
	}
	args := []eliotlibs.PagginationParam{}
	preloads := []eliotlibs.PreloadParams{}
	err = db.Scopes(repo.paginate_with_param(ctx, result, page, args, preloads)).Find(&result).Error
	if err != nil {
		log.Printf("ERROR GET LOFIN FAILEDS %s", err.Error())
		return nil, err
	}
	return result, nil
}
