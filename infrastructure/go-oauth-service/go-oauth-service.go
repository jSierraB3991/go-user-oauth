package gooauthservice

import (
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrepository "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-repository"
	"gorm.io/gorm"
)

type GoOauthService struct {
	repo                    *gooauthrepository.Repository
	passwordService         *PasswordService
	secretForJwt            string
	aesKeyForEncrypt        string
	appName                 string
	urlImagenApp            string
	timeToExpiredQrForOauth time.Duration
}

func NewGoOauthService(database *gorm.DB, secretForJwt, aesKeyForEncrypt string, serviceModel gooauthmodel.ServiceModeParam, timeToExpiredQrForOauth time.Duration) *GoOauthService {
	repo := gooauthrepository.InitiateRepo(database)
	err := repo.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}

	appName := serviceModel.AppName
	if strings.TrimSpace(appName) == "" {
		appName = "Mi APP"
	}

	return &GoOauthService{
		repo:                    repo,
		passwordService:         NewPasswordService(),
		secretForJwt:            secretForJwt,
		aesKeyForEncrypt:        aesKeyForEncrypt,
		appName:                 appName,
		urlImagenApp:            serviceModel.UrlImagenApp,
		timeToExpiredQrForOauth: timeToExpiredQrForOauth,
	}
}

func (GoOauthService) ErrorHandler() error {
	return gooautherror.InactiveToken{}
}

func (s *GoOauthService) GetJwtToken(exp int, userId uint, email string) (string, error) {
	pathsAllow, err := s.repo.GetPathAllowByUser(userId)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"user_id":     userId,
		"exp":         exp,
		"paths_allow": pathsAllow,
		"iat":         time.Now().Unix(),
		"email":       email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretForJwt))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *GoOauthService) GetExp() int {
	return int(time.Now().Add(24 * time.Hour).Unix())
}
