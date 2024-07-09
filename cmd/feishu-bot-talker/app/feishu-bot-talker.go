package app

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func NewCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "feishu-bot-talker",
		Long:    "feishu-bot-talker is cli tool to send message to feishu bot",
		Example: figure.NewColorFigure("FeiShu-Bot-talker", "isometric1", "green", true).String(),
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Print version and exit",
		Long:    "version subcommand will print version and exit",
		Example: "feishu-bot-talker version",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("version:", version)
		},
	}

	cmd.AddCommand(CreateTemplateCommand())
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(CreateSendCommand())
	return cmd
}
