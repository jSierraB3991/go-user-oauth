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

func (repo *Repository) GetDataLoginSessions(ctx context.Context, page *eliotlibs.Paggination) ([]gooauthmodel.GoUserDataLogin, error) {
	var result []gooauthmodel.GoUserDataLogin
	db, err := repo.WithTenant(ctx)
	if err != nil {
		log.Printf("ERROR GET DB %s", err.Error())
		return nil, err
	}
	args := []eliotlibs.PagginationParam{
		{Where: "user_id = ?", Data: []interface{}{page.Data}},
	}
	preloads := []eliotlibs.PreloadParams{}
	err = db.Scopes(repo.paginate_with_param(ctx, result, page, args, preloads)).Find(&result).Error
	if err != nil {
		log.Printf("ERROR GET LOGIN SESSIONS %s", err.Error())
		return nil, err
	}
	return result, nil
}

func (repo *Repository) GetSessionsByToken(ctx context.Context, email, token string) ([]gooauthmodel.GoUserDataLogin, error) {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	user := db.Select("id").Model(&gooauthmodel.GoUserUser{}).Where("email = ?", email)

	var result []gooauthmodel.GoUserDataLogin
	err = db.Where("is_available = ? AND user_id IN ?", true, user).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Repository) RemoveSessionById(ctx context.Context, idSession uint) error {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}
	return db.Model(&gooauthmodel.GoUserUser{}).Where("id = ?", idSession).Update("is_active_two_factor", false).Error
}
