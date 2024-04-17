// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package g

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	cfg *Config
)

type Config struct {
	AppID     string // 应用ID
	AppSecret string // 应用密钥

	// Redis配置
	Redis struct {
		Addr     string // 地址
		Password string // 密码
		DB       int    // 数据库
	}
}

func GetConfig() *Config {
	return cfg
}

func readConfig() (err error) {
	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("配置读取失败: %s\n", err)
	}

	cfg = &Config{}
	err = viper.Unmarshal(cfg)
	return
}

// LoadConfig 加载配置文件
func LoadConfig(configFile string) {
	// 判定配置文件是否存在
	_, err := os.Stat(configFile)
	if err != nil {
		fmt.Println("配置文件不存在")
		os.Exit(1)
	}

	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	// 读取配置文件
	err = readConfig()
	if err != nil {
		fmt.Printf("配置读取失败: %v\n", err)
		os.Exit(1)
	}
}
