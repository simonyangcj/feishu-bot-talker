package main

import (
	"os"

	"feishu-bot-talker/cmd/feishu-bot-talker/app"
)

var version string

func main() {
	cmd := app.NewCommand(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
