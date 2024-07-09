package app

import (
	"encoding/json"
	"feishu-bot-talker/cmd/feishu-bot-talker/app/option"
	api "feishu-bot-talker/pkg/feishu-api"
	"os"

	feiShuUtil "feishu-bot-talker/util/feishu"

	"github.com/spf13/cobra"
)

func CreateTemplateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "handle template",
		Long:  "handle template",
		Example: `feishu-bot-talker template show simple|complex|@someone|@all
		feishu-bot-talker template create simple|complex|@someone|@all`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	option := &option.OptionTemplate{}
	option.BindFlags(cmd.Flags())
	cmd.AddCommand(createShowCommand())
	cmd.AddCommand(createNewCommand())
	return cmd
}

func createNewCommand() *cobra.Command {
	option := &option.OptionTemplate{}
	cmd := &cobra.Command{
		Use:     "new",
		Short:   "create a template to a file",
		Long:    "create a template to a file",
		Example: "feishu-bot-talker template create simple|complex|@someone|@all",
		Run:     createNewCommandHandler(option),
		Args:    cobra.ExactArgs(1),
	}
	option.BindFlags(cmd.Flags())

	return cmd
}

func createShowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "show",
		Short:   "show different template",
		Long:    "show different template",
		Example: "feishu-bot-talker template show simple|complex|@someone|@all",
		Run:     createShowCommandHandler,
		Args:    cobra.ExactArgs(1),
	}

	option := &option.OptionTemplate{}
	option.BindFlags(cmd.Flags())
	return cmd
}

func createShowCommandHandler(cmd *cobra.Command, args []string) {
	if !checkSupportedTemplate(args[0]) {
		cmd.PrintErrf("unsupported template: %s\n", args[0])
		return
	}

	showTemplate := &api.ContentRenderText{}

	switch args[0] {
	case "simple":
		showTemplate.Text = "  test content"
	case "complex":
		showTemplate.Text = " test content \n test content 2 \n test content 3"
	case "@someone":
		showTemplate.Text = "<at user_id=\"ou_xxxxxxx\"></at> hi there"
	case "@all":
		showTemplate.Text = "<at user_id=\"all\"></at> hi all"
	}

	result, err := json.Marshal(showTemplate)
	if err != nil {
		cmd.PrintErrf("failed to marshal template: %v\n", err)
		return
	}

	cmd.Println("Do not remove the blank space in the front of the content !!!!!")
	if args[0] == "@someone" || args[0] == "@all" {
		cmd.Println(feiShuUtil.ReplaceAtTag(string(result)))
		return
	}
	cmd.Println(string(result))
}

func checkSupportedTemplate(template string) bool {
	supportedTemplate := map[string]struct{}{
		"simple":   {},
		"complex":  {},
		"@someone": {},
		"@all":     {},
	}

	_, ok := supportedTemplate[template]
	return ok
}

func createNewCommandHandler(option *option.OptionTemplate) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if !checkSupportedTemplate(args[0]) {
			cmd.PrintErrf("unsupported template: %s\n", args[0])
			return
		}

		showTemplate := &api.ContentRenderText{}

		switch args[0] {
		case "simple":
			showTemplate.Text = "  test content"
		case "complex":
			showTemplate.Text = " test content \n test content 2 \n test content 3"
		case "@someone":
			showTemplate.Text = "<at user_id=\"ou_xxxxxxx\"></at> hi there"
		case "@all":
			showTemplate.Text = "<at user_id=\"all\"></at> hi all"
		}

		result, err := json.Marshal(showTemplate)
		if err != nil {
			cmd.PrintErrf("failed to marshal template: %v\n", err)
			return
		}

		toWrite := result

		if args[0] == "@someone" || args[0] == "@all" {
			toWrite = []byte(feiShuUtil.ReplaceAtTag(string(result)))
		}

		// open io write to file
		err = os.WriteFile(option.OutputFile, toWrite, 0644)
		if err != nil {
			cmd.PrintErrf("failed to write to file: %v\n", err)
			return
		}

		cmd.Println("Do not remove the blank space in the front of the content !!!!!")
		cmd.Printf("write to file: %s\n", option.OutputFile)
		if args[0] == "@someone" {
			cmd.Println("please replace ou_xxxxxxx with real user_id")
		}
	}
}
