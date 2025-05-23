package gooauthrepository

import (
	"context"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
)

func (repo *Repository) SavePath(ctx context.Context, path, operation string) (uint, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return 0, err
	}

	data, err := repo.getPath(ctx, path, operation)
	if err != nil {
		return 0, err
	}

	if data.PathRoute != "" {
		return data.PathBackId, nil
	}

	modelDb := &gooauthmodel.GoUserPathBack{
		PathRoute:      path,
		OperationRoute: operation,
	}
	err = db.Save(&modelDb).Error
	if err != nil {
		return 0, err
	}
	return modelDb.PathBackId, nil
}

func (repo *Repository) SavePathRole(ctx context.Context, pathId uint, roleName string) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	role, err := repo.GetRoleByName(ctx, roleName)
	if err != nil {
		return err
	}

	preData, err := repo.getRolePath(ctx, pathId, role.RoleId)
	if err != nil {
		return err
	}

	if preData.GoUserPathBackId != pathId || preData.GoUserRoleId != role.RoleId {
		return db.Save(&gooauthmodel.GoUserRolePath{GoUserPathBackId: pathId, GoUserRoleId: role.RoleId}).Error
	}
	return nil
}

func (repo *Repository) getRolePath(ctx context.Context, pathId, roleId uint) (*gooauthmodel.GoUserRolePath, error) {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserRolePath
	err = db.Where("role_id = ? AND path_back_id = ?", roleId, pathId).Preload("GoUserRole").Preload("GoUserPathBack").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *Repository) getPath(ctx context.Context, path, operation string) (*gooauthmodel.GoUserPathBack, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserPathBack
	err = db.Where("path_route = ? AND operation_route = ?", path, operation).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
