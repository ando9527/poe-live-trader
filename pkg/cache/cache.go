package cache

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Client struct{
	cache map[string]bool
	sync.Mutex
}

func NewClient() *Client {
	return &Client{
		cache: map[string]bool{},
		Mutex: sync.Mutex{},
	}
}

func (c *Client) AllowSend(id string) bool {
	logrus.Debug("checking duplicated cache")
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if !c.cache[id]{
		c.cache[id]=true
		return true
	}
	logrus.Debug("duplicated user in short time. skipped it.")
	return false
}

func (c *Client) Run() {
	go func() {
		ticker:=time.NewTicker(time.Minute*10)
		for _= range ticker.C{
			c.Mutex.Lock()
			c.cache=make(map[string]bool)
			c.Mutex.Unlock()
		}
	}()
}
