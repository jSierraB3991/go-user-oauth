package gooauthrepository

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
)

func (repo *Repository) SaveUser(user *gooauthmodel.GoUserUser) error {
	var userExist gooauthmodel.GoUserUser
	err := repo.db.Where("email = ?", user.Email).Find(&userExist).Error
	if err != nil {
		return err
	}
	if userExist.Email != "" {
		return gooautherror.UserExistsError{}
	}
	return repo.db.Save(&user).Error
}

func (repo *Repository) SaveAttributtes(userId uint, attr []gooauthmodel.GoUserUserAttributtes) error {
	for i := range attr {
		attr[i].GoUserUserId = userId
	}
	return repo.db.Save(&attr).Error
}

func (repo *Repository) GetUserByEmail(userEmail string) (*gooauthmodel.GoUserUser, error) {
	var result gooauthmodel.GoUserUser
	err := repo.db.Preload("GoUserRole").Where("email = ?", userEmail).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidUserOrPassword{}
	}
	return &result, nil
}

func (repo *Repository) GetUserById(userId uint) (*gooauthmodel.GoUserUser, error) {
	var result gooauthmodel.GoUserUser
	err := repo.db.Preload("GoUserRole").Find(&result, userId).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidUserOrPassword{}
	}
	return &result, nil
}

func (repo *Repository) GetAttributtesByUserId(userId uint) ([]gooauthmodel.GoUserUserAttributtes, error) {
	var result []gooauthmodel.GoUserUserAttributtes
	err := repo.db.Where("user_id = ?", userId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
