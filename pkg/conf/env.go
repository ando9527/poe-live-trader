package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var (
	Config Configuration
)

func init() {
	Config = NewConfig()
}

type Configuration struct {
	League    string `required:"true" split_words:"true"`
	Poesessid string `required:"true" split_words:"true"`
	Filter    string `required:"true" split_words:"true"`
}

// NewConfig Initialize Configuration
func NewConfig() (c Configuration) {
	err := envconfig.Process("app", &c)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return c
}
