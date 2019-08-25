package conf

import (
	"github.com/kelseyhightower/envconfig"
)

var (
	Env EnvironmentVariable
)

type EnvironmentVariable struct {
	League    string  `required:"true" split_words:"true"`
	Poesessid string  `required:"true" split_words:"true"`
	Filter    string  `required:"true" split_words:"true"`
	Volume    float64 `required:"true" split_words:"true"`
	CloudEnable bool `required:"true" split_words:"true"`
}

func InitConfig() (err error) {
	err = envconfig.Process("app", &Env)
	if err != nil {
		return err
	}
	return nil
}
