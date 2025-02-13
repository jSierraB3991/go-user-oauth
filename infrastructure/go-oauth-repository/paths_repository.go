package gooauthrepository

import gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"

func (repo *Repository) SavePath(path, operation string) error {
	modelDb := gooauthmodel.GoUserPathBack{
		PathRoute:      path,
		OperationRoute: operation,
	}

	return repo.db.Save(&modelDb).Error
}
