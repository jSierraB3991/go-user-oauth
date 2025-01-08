package gooauthservice

import (
	"context"

	gooauthmapper "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-mapper"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) UpdateUser(ctx context.Context, keyCloakUserId string, attributes *map[string][]string, req gooauthrequest.UpdateUserRequest) error {
	data, err := s.repo.GetUserById(jsierralibs.GetUNumberForString(keyCloakUserId))
	if err != nil {
		return err
	}

	data.Name = req.FirstName
	data.SubName = req.LastName

	err = s.repo.SaveUser(data)
	if err != nil {
		return err
	}

	attrData, err := s.repo.GetAttributtesByUserId(data.UserId)
	if err != nil {
		return err
	}
	attrDataNews := gooauthmapper.GetAttributtes(attributes)

	finalAttr := gooauthmapper.GetAttributteUpdate(attrData, attrDataNews)

	return s.repo.UpdateAttrr(finalAttr)
}
