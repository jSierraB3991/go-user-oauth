package gooauthservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) saveDataLogin(ctx context.Context, ip, userAgent, token string, userId uint, isLoginWithPassword bool) error {
	timestamp := time.Now().UTC()
	location, err := getIPLocation(ip)
	if err != nil {
		location = &gooauthrequest.IPInfoRequest{Country: "Desconocido", City: "Desconocido"}
	}

	var ipInfo string
	if location != nil {
		locationByte, err := json.Marshal(location)
		if err == nil {
			ipInfo = string(locationByte)
		}
	}

	ipEncrypt, err := jsierralibs.Encrypt(ip, s.aesKeyForEncrypt)
	if err != nil {
		return err
	}

	request := gooauthmodel.GoUserDataLogin{
		Ip:                  ipEncrypt,
		UserAgent:           userAgent,
		IsLoginWithPassword: isLoginWithPassword,
		Token:               token,
		Fecha:               timestamp,
		IpResponse:          ipInfo,
		GoUserUserId:        userId,
		IsAvailable:         true,
	}
	return s.repo.SaveDataLogin(ctx, request)
}

func (s *GoOauthService) saveInvalidDataLogin(ctx context.Context, ip, userAgent, userEmail, motive string, isTwoFactor bool) error {
	timestamp := time.Now().UTC()
	location, err := getIPLocation(ip)
	if err != nil {
		location = &gooauthrequest.IPInfoRequest{Country: "Desconocido", City: "Desconocido"}
	}

	var ipInfo string
	if location != nil {
		locationByte, err := json.Marshal(location)
		if err == nil {
			ipInfo = string(locationByte)
		}
	}

	ipEncrypt, err := jsierralibs.Encrypt(ip, s.aesKeyForEncrypt)
	if err != nil {
		return err
	}

	emailencrypt := userEmail
	if jsierralibs.RemoveSpace(userEmail) != "" {
		emailEncryptPoint, err := jsierralibs.Encrypt(userEmail, s.aesKeyForEncrypt)
		if err != nil {
			return err
		}
		emailencrypt = emailEncryptPoint
	}
	motiveencrypt := motive
	if jsierralibs.RemoveSpace(motive) != "" {
		motiveencryptPoint, err := jsierralibs.Encrypt(motive, s.aesKeyForEncrypt)
		if err != nil {
			return err
		}
		motiveencrypt = motiveencryptPoint
	}

	tenant, _ := jsierralibs.FromContext(ctx)
	request := gooauthmodel.GoUserInvalidGoAuth{
		Ip:                  ipEncrypt,
		UserAgent:           userAgent,
		IsForTwoFactorOauth: isTwoFactor,
		Motive:              motiveencrypt,
		Fecha:               timestamp,
		IpResponse:          ipInfo,
		Email:               emailencrypt,
		TenantId:            tenant,
		IsUtil:              true,
	}
	return s.repo.SaveInvalidLogin(ctx, request)
}

func getIPLocation(ip string) (*gooauthrequest.IPInfoRequest, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var info gooauthrequest.IPInfoRequest
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
