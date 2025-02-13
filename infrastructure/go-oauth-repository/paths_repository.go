package gooauthrepository

import gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"

func (repo *Repository) SavePath(path, operation string) error {
	data, err := repo.getPath(path, operation)
	if err != nil {
		return err
	}

	if data.PathRoute != "" {
		return nil
	}
	modelDb := gooauthmodel.GoUserPathBack{
		PathRoute:      path,
		OperationRoute: operation,
	}

	return repo.db.Save(&modelDb).Error
}

func (repo *Repository) getPath(path, operation string) (*gooauthmodel.GoUserPathBack, error) {
	var result gooauthmodel.GoUserPathBack
	err := repo.db.Where("path_route = ? AND operation_route = ?", path, operation).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
