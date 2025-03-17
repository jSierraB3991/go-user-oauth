package gooauthrepository

import gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"

func (repo *Repository) SavePath(modelDb *gooauthmodel.GoUserPathBack) error {
	data, err := repo.getPath(modelDb.PathRoute, modelDb.OperationRoute)
	if err != nil {
		return err
	}

	if data.PathRoute != "" {
		return nil
	}

	return repo.db.Save(&modelDb).Error
}

func (repo *Repository) SavePathRole(pathId uint, roleName string) error {

	role, err := repo.GetRoleByName(roleName)
	if err != nil {
		return err
	}

	preData, err := repo.getRolePath(pathId, role.RoleId)
	if err != nil {
		return err
	}

	if preData.GoUserPathBackId != pathId || preData.GoUserRoleId != role.RoleId {
		return repo.db.Save(&gooauthmodel.GoUserRolePath{RolePathId: pathId, GoUserRoleId: role.RoleId}).Error
	}
	return nil
}

func (repo *Repository) getRolePath(pathId, roleId uint) (*gooauthmodel.GoUserRolePath, error) {
	var result gooauthmodel.GoUserRolePath
	err := repo.db.Where("role_id = ? AND path_back_id = ?", roleId, pathId).Preload("GoUserRole").Preload("GoUserPathBack").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *Repository) getPath(path, operation string) (*gooauthmodel.GoUserPathBack, error) {
	var result gooauthmodel.GoUserPathBack
	err := repo.db.Where("path_route = ? AND operation_route = ?", path, operation).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
