package gooauthservice

import (
	"log"
	"strings"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrepository "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-repository"
	"gorm.io/gorm"
)

type GoOauthService struct {
	repo             *gooauthrepository.Repository
	passwordService  *PasswordService
	secretForJwt     string
	aesKeyForEncrypt string
	appName          string
	urlImagenApp     string
}

func NewGoOauthService(database *gorm.DB, secretForJwt, aesKeyForEncrypt string, serviceModel gooauthmodel.ServiceModeParam) *GoOauthService {
	repo := gooauthrepository.InitiateRepo(database)
	err := repo.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}

	appName := serviceModel.AppName
	if strings.Trim(appName, " ") == "" {
		appName = "Mi APP"
	}

	return &GoOauthService{
		repo:             repo,
		passwordService:  NewPasswordService(),
		secretForJwt:     secretForJwt,
		aesKeyForEncrypt: aesKeyForEncrypt,
		appName:          appName,
		urlImagenApp:     serviceModel.UrlImagenApp,
	}
}

func (GoOauthService) ErrorHandler() error {
	return gooautherror.InactiveToken{}
}
