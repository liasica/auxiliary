// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package internal

import (
	"github.com/spf13/cobra"

	"auxiliary/internal/g"
)

func RunCommand() {
	var configFile string

	cmd := &cobra.Command{
		Use:               "",
		Short:             "辅助工具控制台",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			// 加载配置文件
			g.LoadConfig(configFile)
		},
	}

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config/config.yaml", "配置文件")

	cmd.AddCommand(
		newMessage(),
	)

	_ = cmd.Execute()
}
