package gooauthrepository

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
)

func (repo *Repository) SaveUser(user *gooauthmodel.User) error {
	return repo.db.Save(&user).Error
}

func (repo *Repository) SaveAttributtes(userId uint, attr []gooauthmodel.UserAttributtes) error {
	for i := range attr {
		attr[i].UserId = userId
	}
	return repo.db.Save(&attr).Error
}

func (repo *Repository) GetUserByEmail(userEmail string) (*gooauthmodel.User, error) {
	var result gooauthmodel.User
	err := repo.db.Where("email = ?", userEmail).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidUserOrPassword{}
	}
	return &result, nil
}

func (repo *Repository) GetUserById(userId uint) (*gooauthmodel.User, error) {
	var result gooauthmodel.User
	err := repo.db.Find(&result, userId).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidUserOrPassword{}
	}
	return &result, nil
}
