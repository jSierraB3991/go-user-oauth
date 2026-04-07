package gooauthservice

import (
	"context"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) UpdateOneAttr(ctx context.Context, keyCloakUserId string, attribute string, value string) error {
	user, err := s.repo.GetUserById(ctx, eliotlibs.GetUNumberForString(keyCloakUserId))
	if err != nil {
		return err
	}

	attrData, err := s.repo.GetAttrByUserId(ctx, user.UserId, attribute)
	if err != nil {
		return err
	}

	attrData.ValueAttributtes = value
	return s.repo.UpdateOneAttr(ctx, *attrData)
}
