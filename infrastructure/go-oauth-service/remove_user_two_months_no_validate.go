package gooauthservice

import "context"

func (s *GoOauthService) RemoveUserTwoMonthsNoValidate(ctx context.Context) error {
	users, err := s.repo.GetUserNoValidateMail(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		err = s.repo.DeleteUser(ctx, user.UserId)
		if err != nil {
			return err
		}
	}
	return nil
}
