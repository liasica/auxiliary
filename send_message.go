// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-17, by liasica

package auxiliary

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

type SendMessageRequest struct {
	ReceiveId string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Content   string `json:"content"`
	UUID      string `json:"uuid"`
}

type InteractiveTemplateMessage struct {
	Type string                         `json:"type"`
	Data InteractiveTemplateMessageData `json:"data"`
}

type InteractiveTemplateMessageData struct {
	TemplateId       string         `json:"template_id"`
	TemplateVariable map[string]any `json:"template_variable"`
}

func NewInteractiveTemplateMessage(templateId string, data map[string]any) *InteractiveTemplateMessage {
	return &InteractiveTemplateMessage{
		Type: "template",
		Data: InteractiveTemplateMessageData{TemplateId: templateId, TemplateVariable: data},
	}
}

func (m *InteractiveTemplateMessage) String() string {
	s, _ := jsoniter.MarshalToString(m)
	return s
}

func (a *App) SendMessage(receiveIdType, receiveId, msgType, content string) (b []byte, data string, err error) {
	var token string
	token, err = a.GetlInternalTenantAccessToken()
	if err != nil {
		return
	}

	b, _ = jsoniter.Marshal(SendMessageRequest{
		ReceiveId: receiveId,
		MsgType:   msgType,
		Content:   content,
		UUID:      uuid.New().String(),
	})

	var res *resty.Response
	res, err = resty.New().R().
		SetQueryParams(map[string]string{
			"receive_id_type": receiveIdType,
		}).
		SetAuthScheme("Bearer").
		SetAuthToken(token).
		SetBody(b).
		Post(UrlSendMessage)

	if res != nil {
		data = res.String()
	}
	return
}
