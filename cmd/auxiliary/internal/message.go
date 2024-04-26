// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package internal

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"

	"auxiliary"
	"auxiliary/internal/g"
)

type message struct {
	*cobra.Command
}

func newMessage() *cobra.Command {
	m := &message{
		&cobra.Command{
			Use:               "message <command>",
			Short:             "消息管理",
			CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		},
	}
	m.addSendCommand()
	return m.Command
}

func (m *message) addSendCommand() {
	type ApkMessage struct {
		ID         string `json:"CI_JOB_ID,omitempty"`
		AppName    string `json:"APP_NAME,omitempty"`
		Message    string `json:"CI_COMMIT_MESSAGE,omitempty"`
		RCIntranet string `json:"RC_INTRANET_DOWNLOAD,omitempty"`
		RCExtranet string `json:"RC_EXTRANET_DOWNLOAD,omitempty"`
		RDIntranet string `json:"RD_INTRANET_DOWNLOAD,omitempty"`
		RDExtranet string `json:"RD_EXTRANET_DOWNLOAD,omitempty"`
		Version    string `json:"VERSION,omitempty"`
	}

	var (
		receiveId  string
		templateId string
		msg        = &ApkMessage{}
	)

	cmd := &cobra.Command{
		Use: "apk [args]",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := g.GetConfig()
			content := auxiliary.NewInteractiveTemplateMessage(templateId, make(map[string]any))
			params, _ := jsoniter.Marshal(msg)
			_ = jsoniter.Unmarshal(params, &content.Data.TemplateVariable)

			b, res, err := auxiliary.NewApp(cfg.AppID, cfg.AppSecret, g.NewRedis()).SendMessage("chat_id", receiveId, "interactive", content.String())
			fmt.Printf("消息请求体: %s\n", string(b))
			if err != nil {
				fmt.Printf("消息请求失败: %v\n", err)
			}
			fmt.Printf("消息发送结果: %s\n", res)
		},
	}

	cmd.Flags().StringVar(&templateId, "template", "", "模板ID")
	cmd.Flags().StringVar(&receiveId, "id", "", "接收消息方")
	cmd.Flags().StringVar(&msg.ID, "job", "", "JOB ID")
	cmd.Flags().StringVar(&msg.RCIntranet, "rc-intranet", "", "RC 内网下载链接")
	cmd.Flags().StringVar(&msg.RCExtranet, "rc-extranet", "", "RC 外网下载链接")
	cmd.Flags().StringVar(&msg.RDIntranet, "rd-intranet", "", "RD 内网下载链接")
	cmd.Flags().StringVar(&msg.RDExtranet, "rd-extranet", "", "RD 外网下载链接")
	cmd.Flags().StringVar(&msg.AppName, "name", "", "App名称")
	cmd.Flags().StringVar(&msg.Message, "message", "", "更新消息")
	cmd.Flags().StringVar(&msg.Version, "version", "", "版本号")

	m.AddCommand(cmd)
}
