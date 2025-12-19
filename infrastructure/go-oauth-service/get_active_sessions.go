package gooauthservice

import (
	"context"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) GetActiveSessions(ctx context.Context, email, tokenString string, page, limit int) (*gooauthrest.LoginSessionRestPagination, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	paggination := &eliotlibs.Paggination{
		Limit: limit,
		Page:  page,
		Data:  user.UserId,
	}
	loginSessions, err := s.repo.GetDataLoginSessions(ctx, paggination)
	if err != nil {
		return nil, err
	}
	return &gooauthrest.LoginSessionRestPagination{
		Limit:      paggination.Limit,
		Page:       paggination.Page,
		TotalRows:  paggination.TotalRows,
		TotalPages: paggination.TotalPages,
		Sort:       paggination.Sort,
		Data:       gooauthmapper.MapLoginSessionsToRest(loginSessions, s.aesKeyForEncrypt, tokenString),
	}, nil
}
