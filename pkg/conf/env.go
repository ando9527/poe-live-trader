package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var (
	Env EnvironmentVariable
)

func init() {
	Env = NewConfig()
}

type EnvironmentVariable struct {
	League    string  `required:"true" split_words:"true"`
	Poesessid string  `required:"true" split_words:"true"`
	Filter    string  `required:"true" split_words:"true"`
	Volume    float64 `required:"true" split_words:"true"`
}

// NewConfig Initialize Configuration
func NewConfig() (c EnvironmentVariable) {
	err := envconfig.Process("app", &c)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return c
}
