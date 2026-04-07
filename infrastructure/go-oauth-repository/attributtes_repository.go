package gooauthrepository

import (
	"context"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
)

func (repo *Repository) GetAttributtesByUserId(ctx context.Context, userId uint) ([]gooauthmodel.GoUserUserAttributtes, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result []gooauthmodel.GoUserUserAttributtes
	err = db.Where("user_id = ?", userId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) SaveAttributtes(ctx context.Context, userId uint, attr []gooauthmodel.GoUserUserAttributtes) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	if attr == nil {
		return nil
	}
	for i := range attr {
		attr[i].GoUserUserId = userId
	}
	return db.Save(&attr).Error
}

func (repo *Repository) UpdateAttrr(ctx context.Context, attr []gooauthmodel.GoUserUserAttributtes) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Save(&attr).Error
}
func (repo *Repository) UpdateOneAttr(ctx context.Context, attr gooauthmodel.GoUserUserAttributtes) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Save(&attr).Error
}

func (repo *Repository) GetAttrByUserId(ctx context.Context, userId uint, attr string) (*gooauthmodel.GoUserUserAttributtes, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserUserAttributtes
	err = db.Where("user_id = ? AND name_attributte = ?", userId, attr).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
