package configuration

import (
	"io"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Trello struct {
			APIKey string `envconfig:"TRELLO_API_KEY" default:""`
			Token  string `envconfig:"TRELLO_TOKEN" default:""`
		}
	}
)

var (
	globalConfig Config
)

func Usage(output io.Writer) {
	if err := envconfig.Usagef("", &globalConfig, output, envconfig.DefaultTableFormat); err != nil {
		panic(err.Error())
	}
}

func Load() {
	envconfig.MustProcess("", &globalConfig)
}

func Get() Config {
	return globalConfig
}