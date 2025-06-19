package gooauthmapper

import (
	"log"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-rest"
)

func GetInvalidLogins(invalidLogins []gooauthmodel.GoUserInvalidGoAuth, decrypt func(encryptedData string, base64Key string) (string, error), aes string) []gooauthrest.InvalidLoginRest {
	var result []gooauthrest.InvalidLoginRest
	for _, v := range invalidLogins {
		emailDecrypt, err := decrypt(v.Email, aes)
		if err != nil {
			log.Printf("INVALID DECRYPT EMAIL %s", err.Error())
		}

		ipDecrypt, err := decrypt(v.Ip, aes)
		if err != nil {
			log.Printf("INVALID DECRYPT IP %s", err.Error())
		}
		motiveDecrypt, err := decrypt(v.Motive, aes)
		if err != nil {
			log.Printf("INVALID DECRYPT MOTIVE %s", err.Error())
		}

		result = append(result, gooauthrest.InvalidLoginRest{
			InvalidId:        v.GoUserInvalidGoAuthId,
			EmailLoginFailed: emailDecrypt,
			IpFailed:         ipDecrypt,
			Motive:           motiveDecrypt,
			Date:             v.Fecha,
			UserAgent:        v.UserAgent,
			IsForTwoFactor:   v.IsForTwoFactorOauth,
		})
	}
	return result
}
