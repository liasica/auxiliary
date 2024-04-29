// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-29, by liasica

package oss

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "oss <command>",
		Short:             "OSS管理",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	cmd.AddCommand(
		uploader(),
	)
	return cmd
}
