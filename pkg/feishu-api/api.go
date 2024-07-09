package feishu_api

import (
	"context"
	"encoding/json"
	"feishu-bot-talker/cmd/feishu-bot-talker/app/option"
	model "feishu-bot-talker/pkg/feishu-model"
	"fmt"
	"os"
	"strings"

	feishuUtil "feishu-bot-talker/util/feishu"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkauth "github.com/larksuite/oapi-sdk-go/v3/service/auth/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func CreateFeiShuDataPost(option *option.Option) (*model.FeiShuDataPost, error) {
	caller := &model.FeiShuDataPost{
		AppID:         option.AppID,
		AppSecret:     option.AppSecret,
		ReceiveIdType: option.ReceiveIdType,
	}
	// read file from option.ContentFile
	content, err := os.ReadFile(option.ContentFile)
	if err != nil {
		return caller, err
	}

	var render ContentRender

	switch option.MessageType {
	case "text":
		render = &ContentRenderText{}
	default:
		return nil, fmt.Errorf("message-type %s is not supported", option.MessageType)
	}

	result, err := render.Render(content)
	if err != nil {
		return caller, err
	}
	caller.Context = string(result)
	caller.Context = feishuUtil.ReplaceAtTag(caller.Context)

	return caller, err
}

func GetTenantAccessToken(option *option.Option) (*model.Token, error) {
	client := lark.NewClient(option.AppID, option.AppSecret)
	req := larkauth.NewInternalTenantAccessTokenReqBuilder().
		Body(larkauth.NewInternalTenantAccessTokenReqBodyBuilder().
			AppId(option.AppID).
			AppSecret(option.AppSecret).
			Build()).
		Build()
	resp, err := client.Auth.TenantAccessToken.Internal(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if !resp.Success() {
		return nil, fmt.Errorf("failed to request token due to error: %s with id: %s", resp.Msg, resp.RequestId())
	}

	token := &model.Token{}
	// parse response body
	err = json.Unmarshal(resp.ApiResp.RawBody, token)
	if err != nil {
		return nil, err
	}

	return token, err
}

func SendMessage(option *option.Option, dataPost *model.FeiShuDataPost) error {
	test := strings.ReplaceAll(dataPost.Context, "\\", "")
	fmt.Println(test)
	client := lark.NewClient(option.AppID, option.AppSecret)
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(option.ReceiveIdType).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(option.ReceiveIdValue).
			MsgType(option.MessageType).
			Content(dataPost.Context).
			Build()).
		Build()

	resp, err := client.Im.Message.Create(context.Background(), req)
	if err != nil {
		return err
	}

	if !resp.Success() {
		return fmt.Errorf("failed to send message due to error: %s with id: %s", resp.Msg, resp.RequestId())
	}

	return nil
}
