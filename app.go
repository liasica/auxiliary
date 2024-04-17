// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package auxiliary

import "github.com/redis/go-redis/v9"

type App struct {
	appID     string
	appSecret string
	cache     *redis.Client
}

func NewApp(appId, appSecret string, rd *redis.Client) *App {
	return &App{
		appID:     appId,
		appSecret: appSecret,
		cache:     rd,
	}
}
