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
	GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error)
	ChangePassword(ctx context.Context, req gooauthrequest.ChangePasswordRequest) error
	ValidateMailByUserId(ctx context.Context, userId string) error

	GenerateQrForDobleOuath(userName string) (*gooauthrest.QrTwoOauthRest, error)
	ValidateCodeOtp(req gooauthrequest.ValidateOauthCodeRequest) (bool, error)

	GeneratetokenToValidate(userId, keyToGenerateToken string, limitInHours time.Duration) (*string, error)
	RemenberPassword(token, newPassword, codeTwoFactor string) error
	IsActiveTwoFactorOauth(token string) (bool, error)

	IsActiveTwoFactor(user string) (bool, error)
	DisAvailableTwoFactorAuth(userEmail, codeTwoFactor string) error
}
