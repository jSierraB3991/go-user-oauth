package gooauthrepository

import (
	"context"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (repo *Repository) SaveUser(ctx context.Context, user *gooauthmodel.GoUserUser) error {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	var userExist gooauthmodel.GoUserUser
	err = db.Where("email = ?", user.Email).Find(&userExist).Error
	if err != nil {
		return err
	}
	if userExist.Email != "" {
		return gooautherror.UserExistsError{}
	}
	return db.Save(&user).Error
}

func (repo *Repository) UpdateUser(ctx context.Context, user *gooauthmodel.GoUserUser) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Save(&user).Error
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

func (repo *Repository) GetUserByEmail(ctx context.Context, userEmail string) (*gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserUser
	err = db.Preload("GoUserRole").Where("email = ?", userEmail).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidUserOrPassword{}
	}
	return &result, nil
}

func (repo *Repository) GetUserById(ctx context.Context, userId uint) (*gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserUser
	err = db.Preload("GoUserRole").Find(&result, userId).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidUserOrPassword{}
	}
	return &result, nil
}

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

func (repo *Repository) GetSecretOauthCode(ctx context.Context, userEmail string) (*string, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserUser
	err = db.Where("email = ?", userEmail).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result.KeyOathApp, nil
}

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

func (repo *Repository) EnableUser(ctx context.Context, userId uint) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Model(&gooauthmodel.GoUserUser{}).Where("id = ?", userId).Update("enabled", true).Error
}

func (repo *Repository) UpdateTokenMailValidatePassword(ctx context.Context, userId uint, tokenString string) error {

	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Model(&gooauthmodel.GoUserUser{}).Where("id = ? ", userId).Update("token_to_change_password", tokenString).Error
}

func (repo *Repository) GetUserByToken(ctx context.Context, token string) (*gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result gooauthmodel.GoUserUser
	err = db.Where("token_to_change_password = ?", token).Find(&result).Error
	if err != nil {
		return nil, err
	}
	if result.Password == "" {
		return nil, gooautherror.InvalidTokenError{}
	}
	return &result, nil
}

func (repo *Repository) UpdateLinkMailValidateMail(ctx context.Context, userId uint, tokenString string) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Model(&gooauthmodel.GoUserUser{}).Where("id = ? ", userId).Update("link_to_validate_mail", tokenString).Error
}

func (repo *Repository) GetUsersByEmail(ctx context.Context, emails []string) ([]gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	if len(emails) == 0 {
		return []gooauthmodel.GoUserUser{}, nil
	}

	var result []gooauthmodel.GoUserUser
	err = db.Where("email IN (?)", emails).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) GetUsersPage(ctx context.Context, page *jsierralibs.Paggination) ([]gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result []gooauthmodel.GoUserUser
	params := []jsierralibs.PagginationParam{}
	preloads := []jsierralibs.PreloadParams{}
	err = db.Scopes(repo.paginate_with_param(ctx, result, page, params, preloads)).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) GetUsersByNamePage(ctx context.Context, page *jsierralibs.Paggination, nameLike string) ([]gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result []gooauthmodel.GoUserUser
	params := []jsierralibs.PagginationParam{{
		Where: "name like ?1 OR sub_name like ?1",
		Data:  []interface{}{"%" + nameLike + "%", "%" + nameLike + "%"},
	}}
	preloads := []jsierralibs.PreloadParams{}
	err = db.Scopes(repo.paginate_with_param(ctx, result, page, params, preloads)).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) ExistsUserAdministrator(ctx context.Context) (bool, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return false, err
	}

	roleAdministrator, err := repo.GetRoleByName(ctx, gooauthlibs.ROLE_ADMIN)
	if err != nil {
		return false, err
	}

	var result gooauthmodel.GoUserUser
	err = db.Where("role_id = ?", roleAdministrator.RoleId).Find(&result).Error
	if err != nil {
		return false, err
	}

	return result.Email != "", nil
}
