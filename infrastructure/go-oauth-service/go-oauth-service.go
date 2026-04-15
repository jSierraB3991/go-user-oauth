package gooauthservice

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthlibs "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_libs"
	gooauthrepository "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-repository"
	"gorm.io/gorm"
)

type GoOauthService struct {
	repo                    *gooauthrepository.Repository
	passwordService         *PasswordService
	secretForJwt            string
	aesKeyForEncrypt        string
	timeToExpiredQrForOauth time.Duration
	genericPasswordForAdmin string
	saveLoginHistory        bool
}

func NewGoOauthService(database *gorm.DB, secretForJwt, aesKeyForEncrypt string,
	timeToExpiredQrForOauth time.Duration, genericPasswordForAdmin string, saveLoginHistory bool) *GoOauthService {
	return NewGoOauthServiceWithSchemas(database,
		secretForJwt,
		aesKeyForEncrypt,
		timeToExpiredQrForOauth,
		[]string{"public"}, genericPasswordForAdmin, saveLoginHistory)
}

func NewGoOauthServiceWithSchemas(database *gorm.DB, secretForJwt, aesKeyForEncrypt string,
	timeToExpiredQrForOauth time.Duration,
	schemas []string, genericPasswordForAdmin string, saveLoginHistory bool) *GoOauthService {
	repo := gooauthrepository.InitiateRepo(database)
	err := repo.RunMigrations(schemas)
	if err != nil {
		log.Fatal(err)
	}

	return &GoOauthService{
		repo:                    repo,
		passwordService:         NewPasswordService(),
		secretForJwt:            secretForJwt,
		aesKeyForEncrypt:        aesKeyForEncrypt,
		timeToExpiredQrForOauth: timeToExpiredQrForOauth,
		genericPasswordForAdmin: genericPasswordForAdmin,
		saveLoginHistory:        saveLoginHistory,
	}
}
func (GoOauthService) ErrorHandler() error {
	return gooautherror.InactiveTokenError{}
}

func (s *GoOauthService) GetJwtToken(ctx context.Context, userId, roleId uint, email, roleName string, remenber bool) (string, int, error) {

	exp := int(0)
	pathsAllow, err := s.repo.GetPathAllowByUser(ctx, userId)
	if err != nil {
		return "", 0, err
	}
	claims := jwt.MapClaims{
		gooauthlibs.USER_ID:        userId,
		gooauthlibs.PATH_ALLOW:     pathsAllow,
		gooauthlibs.IAT:            time.Now().Unix(),
		gooauthlibs.EMAIL_CONSTANT: email,
		gooauthlibs.ROLE_NAME:      roleName,
		gooauthlibs.ROLE_ID:        roleId,
	}
	if !remenber {
		exp = GetExp()
		claims[gooauthlibs.EXP] = exp
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretForJwt))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}

func GetExp() int {
	return int(time.Now().Add(24 * time.Hour).Unix())
}
