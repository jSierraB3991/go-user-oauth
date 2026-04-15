package gooauthservice

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *GoOauthService) saveDataLogin(ctx context.Context, ip, userAgent, refreshToken string, userId uint, isLoginWithPassword bool) error {
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

	ipEncrypt, err := eliotlibs.Encrypt(ip, s.aesKeyForEncrypt)
	if err != nil {
		return err
	}
	refreshTokenEncrypt, err := eliotlibs.Encrypt(refreshToken, s.aesKeyForEncrypt)
	if err != nil {
		return err
	}

	request := gooauthmodel.GoUserDataLogin{
		Ip:                  ipEncrypt,
		UserAgent:           userAgent,
		IsLoginWithPassword: isLoginWithPassword,
		RefreshToken:        &refreshTokenEncrypt,
		Fecha:               timestamp,
		IpResponse:          ipInfo,
		GoUserUserId:        userId,
		IsAvailable:         true,
	}
	return s.repo.SaveDataLogin(ctx, request)
}

func (s *GoOauthService) saveInvalidDataLogin(ctx context.Context, ip, userAgent, userEmail, motive string, isTwoFactor bool) {
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

	ipEncrypt, err := eliotlibs.Encrypt(ip, s.aesKeyForEncrypt)
	if err != nil {
		log.Println(err)
		return
	}

	emailencrypt := userEmail
	if eliotlibs.RemoveSpace(userEmail) != "" {
		emailEncryptPoint, err := eliotlibs.Encrypt(userEmail, s.aesKeyForEncrypt)
		if err != nil {
			log.Println(err)
			return
		}
		emailencrypt = emailEncryptPoint
	}
	motiveencrypt := motive
	if eliotlibs.RemoveSpace(motive) != "" {
		motiveencryptPoint, err := eliotlibs.Encrypt(motive, s.aesKeyForEncrypt)
		if err != nil {
			log.Println(err)
			return
		}
		motiveencrypt = motiveencryptPoint
	}

	tenant, _ := eliotlibs.FromContext(ctx)
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
	s.repo.SaveInvalidLogin(ctx, request)
}

func getIPLocation(ipStr string) (*gooauthrequest.IPInfoRequest, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, errors.New("invalid ip")
	}

	u := url.URL{
		Scheme: "http",
		Host:   "ip-api.com",
		Path:   "/json/" + ip.String(),
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(u.String())
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
