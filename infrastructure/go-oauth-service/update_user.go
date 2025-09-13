package gooauthservice

import (
	"context"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) UpdateUser(ctx context.Context, keyCloakUserId string, attributes *map[string][]string, req gooauthrequest.UpdateUserRequest) error {
	data, err := s.repo.GetUserById(ctx, eliotlibs.GetUNumberForString(keyCloakUserId))
	if err != nil {
		return err
	}

	data.Name = eliotlibs.CapitalizeName(req.FirstName)
	data.SubName = eliotlibs.CapitalizeName(req.LastName)

	err = s.repo.UpdateUser(ctx, data)
	if err != nil {
		return err
	}

	attrData, err := s.repo.GetAttributtesByUserId(ctx, data.UserId)
	if err != nil {
		return err
	}
	attrDataNews := gooauthmapper.GetAttributtes(attributes)

	finalAttr := gooauthmapper.GetAttributteUpdate(attrData, attrDataNews, data.UserId)

	return s.repo.UpdateAttrr(ctx, finalAttr)
}
