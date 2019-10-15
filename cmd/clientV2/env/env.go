package env

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)



type Client struct {
	LogLevel      string  `required:"true" split_words:"true"`
	League      string  `required:"true" split_words:"true"`
	Poesessid   string  `required:"true" split_words:"true"`
	Filter      []string  `required:"true" split_words:"true"`
	Volume      float64 `required:"true" split_words:"true"`
}

func NewClient()*Client  {
	return &Client{
	}
}

func (c *Client)Init(){
	err := envconfig.Process("client", &c)
	if err != nil {
		logrus.Fatal("setup env config properly, " ,err)
	}
}
