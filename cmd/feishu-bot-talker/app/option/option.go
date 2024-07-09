package option

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Option struct {
	AppID          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
	ReceiveIdType  string `json:"receive_id_type"`
	ReceiveIdValue string `json:"receive_id_value"`
	MessageType    string `json:"message_type"`
	ContentFile    string `json:"content_file"`
}

type OptionTemplate struct {
	MessageType string `json:"message_type"`
	OutputFile  string `json:"output_file"`
}

func (option *Option) BindFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&option.AppID, "app-id", "a", "", "App ID")
	fs.StringVarP(&option.AppSecret, "app-secret", "s", "", "App Secret")
	fs.StringVarP(&option.ReceiveIdType, "receive-id-type", "r", "open_id", "Receive ID Type")
	fs.StringVarP(&option.MessageType, "message-type", "m", "text", "Message Type")
	fs.StringVarP(&option.ContentFile, "content-file", "f", "", "content-file")
	fs.StringVarP(&option.ReceiveIdValue, "receive-id-value", "v", "", "receive-id-value")
}

func (option *OptionTemplate) BindFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&option.MessageType, "message-type", "m", "text", "Message Type")
	fs.StringVarP(&option.OutputFile, "output-file", "o", "feishu-template-output", "output-file")
}

func (option *Option) ValidateOptions() error {
	if option.AppID == "" {
		return fmt.Errorf("app-id is required")
	}
	if option.AppSecret == "" {
		return fmt.Errorf("app-secret is required")
	}

	if option.ContentFile == "" {
		return fmt.Errorf("content-file is required")
	}

	if option.ReceiveIdValue == "" {
		return fmt.Errorf("receive-id-value is required")
	}

	supportedReceiveIdType := map[string]struct{}{
		"open_id":  {},
		"union_id": {},
		"chat_id":  {},
		"email":    {},
		"user_id":  {},
	}

	if _, found := supportedReceiveIdType[option.ReceiveIdType]; !found {
		return fmt.Errorf("receive-id-type %s is not supported", option.ReceiveIdType)
	}

	// so far we only support text
	// more supported message type can be added here
	supportedMessageType := map[string]struct{}{
		"text": {},
	}

	if _, found := supportedMessageType[option.MessageType]; !found {
		return fmt.Errorf("message-type %s is not supported", option.MessageType)
	}

	// check ContentFile exists
	if _, err := os.Stat(option.ContentFile); os.IsNotExist(err) {
		return fmt.Errorf("content-file %s does not exist", option.ContentFile)
	}

	return nil
}

func (option *OptionTemplate) ValidateOptions() error {
	supportedMessageType := map[string]struct{}{
		"text": {},
	}

	if _, found := supportedMessageType[option.MessageType]; !found {
		return fmt.Errorf("message-type %s is not supported", option.MessageType)
	}

	if option.OutputFile == "" {
		return fmt.Errorf("output-file is required")
	}

	return nil
}
