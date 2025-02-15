package gooauthservice

import (
	"log"

	gooautherror "github.com/jSierraB3991/go-user-oauth/domain/go_oauth_error"
	gooauthrepository "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-repository"
	"gorm.io/gorm"
)

type GoOauthService struct {
	repo             *gooauthrepository.Repository
	passwordService  *PasswordService
	secretForJwt     string
	aesKeyForEncrypt string
}

func NewGoOauthService(database *gorm.DB, secretForJwt, aesKeyForEncrypt string) *GoOauthService {
	repo := gooauthrepository.InitiateRepo(database)
	err := repo.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}

	return &GoOauthService{
		repo:             repo,
		passwordService:  NewPasswordService(),
		secretForJwt:     secretForJwt,
		aesKeyForEncrypt: aesKeyForEncrypt,
	}
}

func (GoOauthService) ErrorHandler() error {
	return gooautherror.InactiveToken{}
}
