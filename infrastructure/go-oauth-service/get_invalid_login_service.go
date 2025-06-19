package gooauthservice

import (
	"context"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) GetInvalidLogins(ctx context.Context, page *eliotlibs.Paggination) (*gooauthrest.InvalidLoginRestPagg, error) {
	data, err := s.repo.GetDataLoginFailed(ctx, page)
	if err != nil {
		return nil, err
	}
	return &gooauthrest.InvalidLoginRestPagg{
		Limit:      page.Limit,
		Page:       page.Page,
		TotalRows:  page.TotalRows,
		TotalPages: page.TotalPages,
		Sort:       page.Sort,
		Data:       gooauthmapper.GetInvalidLogins(data, eliotlibs.Decrypt, s.aesKeyForEncrypt),
	}, nil
}
