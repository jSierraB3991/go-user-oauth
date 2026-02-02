package gooauthservice

import (
	"context"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) RemoveUserTwoMonthsNoValidate(ctx context.Context) ([]string, error) {
	users, err := s.repo.GetUserNoValidateMail(ctx)
	if err != nil {
		return nil, err
	}
	var deletedUsers []string
	for _, user := range users {
		err = s.repo.DeleteUser(ctx, user.UserId)
		if err != nil {
			return deletedUsers, err
		}
		deletedUsers = append(deletedUsers, eliotlibs.GetStringUNumberFor(user.UserId))
	}
	return deletedUsers, nil
}
