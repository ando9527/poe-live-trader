package env

import (
	"github.com/kelseyhightower/envconfig"
)



type Env struct {
	LogLevel      string  `required:"true" split_words:"true"`
	League      string  `required:"true" split_words:"true"`
	Poesessid   string  `required:"true" split_words:"true"`
	Filter      []string  `required:"true" split_words:"true"`
	Volume      float64 `required:"true" split_words:"true"`
}

func NewEnv()(cfg Env, err error)  {
	cfg = Env{}
	err = envconfig.Process("client", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
