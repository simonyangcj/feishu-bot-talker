package option

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Option struct {
	ConfigFile string `json:"config_file" yaml:"configFile"`
}

func (option *Option) BindFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&option.ConfigFile, "config-file", "c", "", "config-file")
}

func (option *Option) ValidateOptions() error {
	if option.ConfigFile == "" {
		return fmt.Errorf("config-file is required")
	}

	if _, err := os.Stat(option.ConfigFile); os.IsNotExist(err) {
		return fmt.Errorf("config file %s does not exist", option.ConfigFile)
	}

	return nil
}
