package gooauthrepository

import (
	"context"
	"encoding/json"
	"time"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	"gorm.io/gorm"
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

func (repo *Repository) GetUsersByIds(ctx context.Context, userIds []uint) ([]gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result []gooauthmodel.GoUserUser
	err = db.Where("id IN ?", userIds).Preload("GoUserRole").Find(&result).Error
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
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

func (repo *Repository) GetUsersPage(ctx context.Context, page *eliotlibs.Paggination) ([]gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var result []gooauthmodel.GoUserUser
	params := []eliotlibs.PagginationParam{}
	preloads := []eliotlibs.PreloadParams{}
	err = db.Scopes(repo.paginate_with_param(ctx, result, page, params, preloads)).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) GetUsersByNamePage(ctx context.Context, page *eliotlibs.Paggination, nameLike string) ([]gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}
	val := "%" + nameLike + "%"
	var result []gooauthmodel.GoUserUser
	params := []eliotlibs.PagginationParam{{
		Where: "name ILIKE $1 OR sub_name ILIKE $1",
		Data:  []interface{}{val},
	}}
	preloads := []eliotlibs.PreloadParams{}
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

func (repo *Repository) VerifyIfEmailInAnotherAccont(ctx context.Context, newEmail string) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	var result gooauthmodel.GoUserUser
	err = db.Preload("GoUserRole").Where("email = ?", newEmail).Find(&result).Error
	if err != nil {
		return err
	}

	if result.Email != "" {
		return gooautherror.NewEmailInAntherUserError{}
	}
	return nil
}

func (repo *Repository) GetUserNoValidateMail(ctx context.Context, usersNoRemove []uint) ([]gooauthmodel.GoUserUser, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	twoMonthsAgo := time.Now().AddDate(0, -2, 0)

	var result []gooauthmodel.GoUserUser
	err = db.
		Where("enabled = ? AND created_at <= ? AND id NOT IN (?)", false, twoMonthsAgo, usersNoRemove).
		Order("created_at DESC").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repository) getDataUserString(ctx context.Context, userId uint) (string, string, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return "", "", err
	}

	var attr []gooauthmodel.GoUserUserAttributtes
	if err := db.
		Where("user_id = ?", userId).
		Find(&attr).Error; err != nil {
		return "", "", err
	}

	var user gooauthmodel.GoUserUser
	if err := db.
		Find(&user, userId).Error; err != nil {
		return "", "", err
	}

	userString, err := json.Marshal(&user)
	if err != nil {
		return "", "", err
	}

	userAttrString, err := json.Marshal(&attr)
	if err != nil {
		return "", "", err
	}

	return string(userString), string(userAttrString), nil
}

func (repo *Repository) DeleteUser(ctx context.Context, userId uint) error {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {

		userString, userAttrString, err := repo.getDataUserString(ctx, userId)
		if err != nil {
			return err
		}

		dataToSave := gooauthmodel.UserDataRemove{
			DataUserPpal: string(userString),
			DataUserAttr: string(userAttrString),
		}

		if err := tx.Create(&dataToSave).Error; err != nil {
			return err
		}

		// borrar atributos del usuario (hard delete)
		if err := tx.
			Unscoped().
			Delete(&gooauthmodel.GoUserUserAttributtes{}, "user_id = ?", userId).
			Error; err != nil {
			return err
		}

		// borrar usuario (hard delete)
		if err := tx.
			Unscoped().
			Delete(&gooauthmodel.GoUserUser{}, userId).
			Error; err != nil {
			return err
		}

		return nil
	})
}
