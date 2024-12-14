package gooauthinterface

import (
	"context"
	"net/http"

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

	GetUserByUserId(ctx context.Context, keycloakId string) (*gooauthrest.User, error)
	GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error)
	ChangePassword(ctx context.Context, keycloakUserId, newPassword string) error
	ValidateMailByUserId(ctx context.Context, userId string) error
}