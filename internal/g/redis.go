// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package g

import "github.com/redis/go-redis/v9"

// NewRedis 初始化redis
func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}
