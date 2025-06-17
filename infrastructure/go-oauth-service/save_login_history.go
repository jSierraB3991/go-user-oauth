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

	request := gooauthmodel.GoUserDataLogin{
		Ip:                  ip,
		UserAgent:           userAgent,
		IsLoginWithPassword: isLoginWithPassword,
		Token:               token,
		Fecha:               timestamp,
		IpResponse:          ipInfo,
		GoUserUserId:        userId,
	}
	return s.repo.SaveDataLogin(ctx, request)
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
