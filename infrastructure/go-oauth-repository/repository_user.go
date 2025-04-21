package gooauthrepository

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
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

func (repo *Repository) UpdateUser(user *gooauthmodel.GoUserUser) error {
	return repo.db.Save(&user).Error
}

func (repo *Repository) SaveAttributtes(userId uint, attr []gooauthmodel.GoUserUserAttributtes) error {
	if attr == nil {
		return nil
	}
	for i := range attr {
		attr[i].GoUserUserId = userId
	}
	return repo.db.Save(&attr).Error
}

func (repo *Repository) UpdateAttrr(attr []gooauthmodel.GoUserUserAttributtes) error {
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

func (repo *Repository) SaveSecretToUser(userEmail, keyOath string) error {
	userData, err := repo.GetUserByEmail(userEmail)
	if err != nil {
		return err
	}

	if userData.IsActiveTwoFactorOauth {
		return gooautherror.InvalidTwoFactorIsActive{}
	}

	userData.KeyOathApp = keyOath
	return repo.db.Save(&userData).Error
}

func (repo *Repository) GetSecretOauthCode(userEmail string) (*string, error) {
	var result gooauthmodel.GoUserUser
	err := repo.db.Where("email = ?", userEmail).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result.KeyOathApp, nil
}

func (repo *Repository) ActiveTwoFactorOauth(userEmail string) error {
	userData, err := repo.GetUserByEmail(userEmail)
	if err != nil {
		return err
	}

	userData.IsActiveTwoFactorOauth = true
	return repo.db.Save(&userData).Error
}

func (repo *Repository) EnableUser(userId uint) error {
	return repo.db.Model(&gooauthmodel.GoUserUser{}).Where("id = ?", userId).Update("enabled", true).Error
}

func (repo *Repository) UpdateTokenMailValidatePassword(userId uint, tokenString string) error {
	return repo.db.Model(&gooauthmodel.GoUserUser{}).Where("id = ? ", userId).Update("token_to_change_password", tokenString).Error
}

func (repo *Repository) GetUserByToken(token string) (*gooauthmodel.GoUserUser, error) {
	var result gooauthmodel.GoUserUser
	err := repo.db.Where("token_to_change_password = ?", token).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidTokenError{}
	}
	return &result, nil
}

func (repo *Repository) UpdateLinkMailValidateMail(userId uint, tokenString string) error {
	return repo.db.Model(&gooauthmodel.GoUserUser{}).Where("id = ? ", userId).Update("link_to_validate_mail", tokenString).Error
}

func (repo *Repository) GetUsersByEmail(emails []string) ([]gooauthmodel.GoUserUser, error) {
	if len(emails) == 0 {
		return []gooauthmodel.GoUserUser{}, nil
	}

	var result []gooauthmodel.GoUserUser
	err := repo.db.Where("email IN (?)", emails).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) GetUsersPage(page *jsierralibs.Paggination) ([]gooauthmodel.GoUserUser, error) {
	var result []gooauthmodel.GoUserUser
	params := []jsierralibs.PagginationParam{}
	preloads := []jsierralibs.PreloadParams{}
	err := repo.db.Scopes(repo.paginate_with_param(result, page, params, preloads)).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) ExistsUserAdministrator() (bool, error) {

	roleAdministrator, err := repo.GetRoleByName(gooauthlibs.ROLE_ADMIN)
	if err != nil {
		return false, err
	}

	var result gooauthmodel.GoUserUser
	err = repo.db.Where("role_id = ?", roleAdministrator.RoleId).Find(&result).Error
	if err != nil {
		return false, err
	}

	return result.Email != "", nil
}
