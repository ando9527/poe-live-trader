package conf

import (
	"github.com/kelseyhightower/envconfig"
)

var (
	Env Config
)

type Config struct {
	LogLevel      string  `required:"true" split_words:"true"`
	League      string  `required:"true" split_words:"true"`
	Poesessid   string  `required:"true" split_words:"true"`
	Filter      string  `required:"true" split_words:"true"`
	Volume      float64 `required:"true" split_words:"true"`
	CloudEnable bool    `required:"true" split_words:"true"`
	CloudUrl    string  `required:"true" split_words:"true"`
	Pass        string	`required:"true" split_words:"true"`
	User        string  `required:"true" split_words:"true"`

}

func InitConfig() (err error) {
	err = envconfig.Process("client", &Env)
	if err != nil {
		return err
	}
	return nil
}
