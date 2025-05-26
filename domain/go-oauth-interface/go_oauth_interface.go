package gooauthinterface

import (
	"context"
	"net/http"
	"time"

	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

type GoOauthInterface interface {
	CheckoutMiddleware(requets *http.Request) bool
	GetSecretByClient(ctx context.Context) error
	CreateUser(ctx context.Context, userParam gooauthrequest.CreateUser, roleUser string, attributes *map[string][]string) (string, error)
	UpdateUser(ctx context.Context, keyCloakUserId string, attributes *map[string][]string, req gooauthrequest.UpdateUserRequest) error
	ErrorHandler() error
	GetUserByRole(ctx context.Context, role string) ([]*gooauthrest.User, error)
	LoginUser(ctx context.Context, userName, password string) (*gooauthrest.JWT, error)
	LoginWithTwoFactor(ctx context.Context, userName, codeTwoFactor string) (*gooauthrest.JWT, error)

	GetUserByUserId(ctx context.Context, keycloakId string) (*gooauthrest.User, error)
	GetUserByEmail(ctx context.Context, email string) (*gooauthrest.User, error)
	GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error)
	ChangePassword(ctx context.Context, req gooauthrequest.ChangePasswordRequest) error
	ValidateMailByUserId(ctx context.Context, userId string) error

	GenerateQrForDobleOuath(ctx context.Context, userName, appName, imageUrl string) (*gooauthrest.QrTwoOauthRest, error)
	ValidateCodeOtp(ctx context.Context, req gooauthrequest.ValidateOauthCodeRequest) (bool, error)

	GeneratetokenToValidate(ctx context.Context, userId, keyToGenerateToken string, limitInHours time.Duration) (*string, error)
	GenerateValidateMail(ctx context.Context, mailSend, keyToGenerateToken string) (string, error)
	RemenberPassword(ctx context.Context, token, newPassword, codeTwoFactor string) error
	IsActiveTwoFactorOauth(ctx context.Context, token string) (bool, error)

	IsActiveTwoFactor(ctx context.Context, user string) (bool, error)
	DisAvailableTwoFactorAuth(ctx context.Context, userEmail, codeTwoFactor string) error
	GetUsersByEmail(ctx context.Context, emails []string) ([]gooauthrest.User, error)

	CreateUserAdministrator(ctx context.Context, userEmail, userpassword, appName string, attributes *map[string][]string) (string, error)
	ExistsUserAdministrator(ctx context.Context) (bool, error)
}
