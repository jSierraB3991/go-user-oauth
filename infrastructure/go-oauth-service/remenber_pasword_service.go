package gooauthservice

func (s *GoOauthService) RemenberPassword(token, newPassword string) error {
	userData, err := s.repo.GetUserByToken(token)
	if err != nil {
		return err
	}

	encryptPasword, err := s.passwordService.EncryptPassword(newPassword)
	if err != nil {
		return err
	}
	userData.Password = encryptPasword
	userData.TokenChangePassword = ""
	return s.repo.UpdateUser(userData)
}

func (s *GoOauthService) ValidateToken(token string) error {
	_, err := s.repo.GetUserByToken(token)
	return err
}
