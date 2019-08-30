package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

func NewConfig()(cfg Config)  {
	cfg = Config{}
	err := envconfig.Process("cloud", &cfg)
	if err != nil {
		logrus.Panic(errors.Wrap(err, "Please setup .env file properly"))
	}
	return cfg
}
