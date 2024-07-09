package app

import (
	"feishu-bot-talker/cmd/feishu-bot-talker/app/option"
	api "feishu-bot-talker/pkg/feishu-api"

	"github.com/spf13/cobra"
)

func CreateSendCommand() *cobra.Command {
	option := &option.Option{}
	cmd := &cobra.Command{
		Use:     "send",
		Short:   "send message to feishu bot with given template",
		Long:    "send message to feishu bot with given template",
		Example: `feishu-bot-talker --app-id=id1 --app-secret=secret1 --receive-id-type=chat_id --receive-id-value=real-chat-id --message-type=text --content-file=template-file send`,
		Run:     CreateSendCommandHandler(option),
	}

	option.BindFlags(cmd.Flags())
	return cmd
}

func CreateSendCommandHandler(option *option.Option) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := option.ValidateOptions()
		if err != nil {
			cmd.PrintErrln(err)
			cmd.Help()
			return
		}

		postData, err := api.CreateFeiShuDataPost(option)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err = api.SendMessage(option, postData)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	}
}
