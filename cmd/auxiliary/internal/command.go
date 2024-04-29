// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package internal

import (
	"fmt"

	"github.com/spf13/cobra"

	"auxiliary/cmd/auxiliary/internal/oss"
	"auxiliary/internal/g"
)

var (
	DefaultConfigPath = "config/config.yaml"
)

func RunCommand() {
	var configFile string

	cmd := &cobra.Command{
		Use:               "",
		Short:             "辅助工具控制台",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			fmt.Printf("使用配置文件: %s\n", configFile)
			// 加载配置文件
			g.LoadConfig(configFile)
		},
	}

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", DefaultConfigPath, "配置文件")

	cmd.AddCommand(
		newMessage(),
		oss.New(),
	)

	_ = cmd.Execute()
}
