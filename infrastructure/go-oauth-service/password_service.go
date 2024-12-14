package gooauthservice

import "golang.org/x/crypto/bcrypt"

type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ps *PasswordService) VerifyPassword(passwordDb, paswordLogin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDb), []byte(paswordLogin))
	return err == nil
}
