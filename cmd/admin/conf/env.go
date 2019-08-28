package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Debug        bool     `required:"true" split_words:"true"`
	GoogleProjectId string `required:"true" split_words:"true"`
	Dsn string `required:"true" split_words:"true"`


}

type Auth struct{
	GoogleApplicationCredentials  string `required:"true" split_words:"true"`
}

func InitAuth(){
	auth:=Auth{}
	err := envconfig.Process("", &auth)
	if err != nil {
		logrus.Panic(errors.Wrap(err, "Please setup .env file properly"))
	}
}



func NewConfig() (cfg Config) {
	InitAuth()
	cfg = Config{}
	err := envconfig.Process("app", &cfg)
	if err != nil {
		logrus.Panic(errors.Wrap(err, "Please setup .env file properly"))
	}
	return cfg
}
