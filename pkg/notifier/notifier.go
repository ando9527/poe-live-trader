package notifier

import (
	"context"

	"github.com/ando9527/poe-live-trader/pkg/key"
	"github.com/sirupsen/logrus"
)

type Client struct{
    queue chan string
    key *key.Client
}

func NewClient(ctx context.Context) *Client {
	return &Client{
		queue: make(chan string),
		key:   key.NewClient(ctx),
	}
}

func (c Client) Run() {
	c.key.Run()
	go func(){
		for v:=range c.queue{
			err := c.key.InsertByRobotGo(v)
			if err != nil {
				logrus.Error(err)
			}
		}
	}()
}

func (c Client) SendToQueue(m string) {
	c.queue<-m
}

