package gooauthrepository

import (
	"context"
	"log"
	"time"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	"gorm.io/gorm"
)

func (repo *Repository) GetSessionsByEmailRefreshTokenE(ctx context.Context, email, refreshTokenE string) ([]gooauthmodel.GoUserDataLogin, error) {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	user := db.Select("id").Model(&gooauthmodel.GoUserUser{}).Where("email = ?", email)

	var result []gooauthmodel.GoUserDataLogin
	err = db.Where("is_available = ? AND user_id IN (?) AND refresh_token = ?", true, user, refreshTokenE).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (repo *Repository) GetSessionsByRefreshToken(ctx context.Context, refreshToken string) (*gooauthmodel.GoUserDataLogin, error) {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserDataLogin
	err = db.Preload("GoUserUser").Preload("GoUserUser.GoUserRole").Where("is_available = ? AND refresh_token = ?", true, refreshToken).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gooautherror.NotFoundSessionByRefreshTokenError{}
		}
		return nil, err
	}

	return &result, nil
}

func (repo *Repository) RemoveSessionById(ctx context.Context, idSession uint) error {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}
	return db.Model(&gooauthmodel.GoUserDataLogin{}).Where("id = ?", idSession).Update("is_available", false).Error
}

func (repo *Repository) GetSessionById(ctx context.Context, idSession uint) (*gooauthmodel.GoUserDataLogin, error) {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}
	var result gooauthmodel.GoUserDataLogin
	err = db.Preload("GoUserUser").Preload("GoUserUser.GoUserRole").Where("id = ?", idSession).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gooautherror.NotFoundSessionError{}
		}
		return nil, err
	}
	return &result, nil
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

func (repo *Repository) RemoveSessionsPreDate(ctx context.Context, limit time.Time) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Model(&gooauthmodel.GoUserDataLogin{}).
		Where("is_available = ?", true).
		Where("updated_at <= ?", limit).
		Update("is_available", false).Error
}

func (repo *Repository) UpdateRefreshToken(ctx context.Context, idSession uint, refreshToken string) error {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Model(&gooauthmodel.GoUserDataLogin{}).
		Where("id = ?", idSession).
		Update("refresh_token = ?", refreshToken).Error
}
