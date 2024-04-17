// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package auxiliary

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

type InternalTenantAccessTokenRequest struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type InternalTenantAccessTokenResponse struct {
	Code              int    `json:"code"`                // 错误码，非 0 取值表示失败
	Msg               string `json:"msg"`                 // 错误描述
	TenantAccessToken string `json:"tenant_access_token"` // 租户访问凭证
	Expire            int    `json:"expire"`              // tenant_access_token 的过期时间，单位为秒
}

const (
	InternalTenantAccessTokenCacheKey = "AUXILIARY:INTERNALTENANTACCESSTOKENCACHEKEY"
)

func (a *App) GetlInternalTenantAccessToken() (token string, err error) {
	token = a.cache.Get(context.Background(), InternalTenantAccessTokenCacheKey).Val()
	if token != "" {
		return
	}
	return a.RequestInternalTenantAccessToken()
}

func (a *App) RequestInternalTenantAccessToken() (token string, err error) {
	var result InternalTenantAccessTokenResponse
	_, err = resty.New().R().
		SetBody(&InternalTenantAccessTokenRequest{
			AppId:     a.appID,
			AppSecret: a.appSecret,
		}).
		SetResult(&result).
		Post(UrlInternalTenantAccessToken)
	if err != nil {
		return
	}

	if result.Code != 0 {
		return "", NewError(result.Code, result.Msg)
	}

	_ = a.cache.Set(context.Background(), InternalTenantAccessTokenCacheKey, result.TenantAccessToken, time.Second*time.Duration(result.Expire-1)).Err()
	token = result.TenantAccessToken
	return
}
