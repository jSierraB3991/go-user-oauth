package gooauthrepository

import (
	"context"
	"log"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (repo *Repository) SaveDataLogin(ctx context.Context, dataLogin gooauthmodel.GoUserDataLogin) (uint, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return 0, err
	}

	err = db.Save(&dataLogin).Error
	if err != nil {
		return 0, err
	}
	return dataLogin.UserDataLoginId, nil
}

func (repo *Repository) SaveInvalidLogin(ctx context.Context, invalidDataLogin gooauthmodel.GoUserInvalidGoAuth) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return
	}

	err = db.Save(&invalidDataLogin).Error
	if err != nil {
		log.Println(err)
	}
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
