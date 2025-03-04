package gooauthrepository

import (
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
)

func (repo *Repository) GetRoleByName(roleName string) (*gooauthmodel.GoUserRole, error) {
	var result gooauthmodel.GoUserRole
	err := repo.db.Where("role_name = ?", roleName).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if result.RoleName == "" {
		return nil, gooautherror.InvalidRole{}
	}
	return &result, nil
}

func (repo *Repository) MigrateO1() error {
	err := repo.db.Save(&gooauthmodel.GoUserRole{RoleName: gooauthlibs.ROLE_USER}).Error
	if err != nil {
		return err
	}
	return repo.db.Save(&gooauthmodel.GoUserRole{RoleName: gooauthlibs.ROLE_ADMIN}).Error
}

func (repo *Repository) GetRolesByUserAndRole(userId, roleId uint) ([]string, error) {
	var result []string

	var rolePath []gooauthmodel.GoUserRolePath
	err := repo.db.Where("role_id = ?", roleId).Preload("GoUserPathBack").Find(&rolePath).Error
	if err != nil {
		return nil, err
	}

	for _, v := range rolePath {
		result = append(result, v.GoUserPathBack.PathRoute)
	}

	var userPath []gooauthmodel.GoUserUserPath
	err = repo.db.Where("user_id = ?", userId).Preload("GoUserPathBack").Find(&userPath).Error
	if err != nil {
		return nil, err
	}

	for _, v := range userPath {
		result = append(result, v.GoUserPathBack.PathRoute)
	}
	return result, nil
}

func (repo *Repository) GetPathAllowByUser(userId uint) ([]string, error) {
	user, err := repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	roles, err := repo.GetRolesByUserAndRole(userId, user.GoUserRoleId)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
