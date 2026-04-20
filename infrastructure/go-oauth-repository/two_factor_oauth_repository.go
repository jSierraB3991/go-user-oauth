package gooauthrepository

import (
	"context"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
)

func (repo *Repository) ActiveTwoFactorOauth(ctx context.Context, userEmail string) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	userData, err := repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	userData.IsActiveTwoFactorOauth = true
	return db.Save(&userData).Error
}

func (repo *Repository) SaveSecretToUser(ctx context.Context, userEmail, keyOath string) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	userData, err := repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	if userData.IsActiveTwoFactorOauth {
		return gooautherror.InvalidTwoFactorIsActive{}
	}

	userData.KeyOathApp = keyOath
	return db.Save(&userData).Error
}

func (repo *Repository) GetUserLoginDataByTokenUUID(ctx context.Context, tokenUuid string) (*gooauthmodel.GoUserDataLogin, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserDataLogin
	err = db.Where("token_two_factor = ?", tokenUuid).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
