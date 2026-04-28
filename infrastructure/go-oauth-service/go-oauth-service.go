package gooauthservice

import (
	"context"
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
	timeToExpiredSession    time.Duration
	genericPasswordForAdmin string
	saveLoginHistory        bool
	schemas                 []string
}

func NewGoOauthService(database *gorm.DB, secretForJwt, aesKeyForEncrypt string,
	timeToExpiredQrForOauth, timeToExpiredSession time.Duration, genericPasswordForAdmin string, saveLoginHistory bool) (*GoOauthService, error) {
	return NewGoOauthServiceWithSchemas(database,
		secretForJwt,
		aesKeyForEncrypt,
		timeToExpiredQrForOauth,
		timeToExpiredSession,
		[]string{"public"}, genericPasswordForAdmin, saveLoginHistory)
}

func NewGoOauthServiceWithSchemas(database *gorm.DB, secretForJwt, aesKeyForEncrypt string,
	timeToExpiredQrForOauthInMinutes, timeToExpiredSessionInMinutes time.Duration,
	schemas []string, genericPasswordForAdmin string, saveLoginHistory bool) (*GoOauthService, error) {

	if timeToExpiredSessionInMinutes > time.Duration(gooauthlibs.MAXIMUM_MINUTES_SESSION)*time.Minute {
		return nil, gooautherror.InvalidSessionTimeError{TimeMinutes: timeToExpiredSessionInMinutes}
	}

	repo := gooauthrepository.InitiateRepo(database)
	err := repo.RunMigrations(schemas)
	if err != nil {
		return nil, err
	}

	return &GoOauthService{
		repo:                    repo,
		passwordService:         NewPasswordService(),
		secretForJwt:            secretForJwt,
		aesKeyForEncrypt:        aesKeyForEncrypt,
		timeToExpiredQrForOauth: time.Duration(timeToExpiredQrForOauthInMinutes) * time.Minute,
		timeToExpiredSession:    time.Duration(timeToExpiredSessionInMinutes) * time.Minute,
		genericPasswordForAdmin: genericPasswordForAdmin,
		saveLoginHistory:        saveLoginHistory,
		schemas:                 schemas,
	}, nil
}
func (GoOauthService) ErrorHandler() error {
	return gooautherror.InactiveTokenError{}
}

func (s *GoOauthService) GetJwtToken(ctx context.Context, userId, roleId uint, email, roleName string, sessionId uint, remenber bool) (string, int, error) {

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
		gooauthlibs.SESSION_ID:     sessionId,
	}
	if !remenber {
		exp = GetExp(s.timeToExpiredSession)
		claims[gooauthlibs.EXP] = exp
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretForJwt))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}

func GetExp(t time.Duration) int {
	return int(time.Now().Add(t).Unix())
}

func (s *GoOauthService) RefreshDatabase(db *gorm.DB) {
	s.repo.SetDb(db)
	s.repo.RunMigrations(s.schemas)
}
