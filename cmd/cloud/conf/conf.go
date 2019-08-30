package conf

import (
"github.com/kelseyhightower/envconfig"
"github.com/pkg/errors"
"github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel        string     `required:"true" split_words:"true"`
	User string   `required:"true" split_words:"true"`
	Pass string   `required:"true" split_words:"true"`
	GoogleProjectId string `required:"true" split_words:"true"`
	Dsn string `required:"true" split_words:"true"`

}



func NewConfig()(cfg Config)  {
	cfg = Config{}
	err := envconfig.Process("cloud", &cfg)
	if err != nil {
		logrus.Panic(errors.Wrap(err, "Please setup .env file properly"))
	}
	return cfg
}
