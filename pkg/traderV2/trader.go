package traderV2

import (
	"github.com/ando9527/poe-live-trader/cmd/clientV2/env"
	"github.com/ando9527/poe-live-trader/pkg/types"
)

type Client struct{
	env *env.Client
	database Database
	wsPool WsPool
	idCache IDCache
	notifier Notifier
	httpClient types.HttpClient
}

func NewClient(cfg *env.Client) *Client {
	return &Client{
		env: cfg,
	}
}

func (c *Client)Run()  {
	c.database.Connect()
	c.database.Migration()
	c.idCache.Run()
	c.wsPool.Run()
	c.notifier.Run()

	for v:=range c.wsPool.GetBuilderChannel(){
		go func(){
			item:=v.SetWhisper(&c.httpClient).Build()
			if allow := c.idCache.AllowSend(item.GetUserID());!allow{
				return
			}

			if ignored:= c.database.isIgnored(item.GetUserID());ignored{
				return
			}

			c.notifier.SendToQueue(item.GetNotification())
		}()
	}

}


