package backend


import (
	"context"

	"github.com/ando9527/poe-live-trader/cmd/client/env"
	"github.com/ando9527/poe-live-trader/pkg/cache"
	"github.com/ando9527/poe-live-trader/pkg/db"
	"github.com/ando9527/poe-live-trader/pkg/notifier"
	"github.com/ando9527/poe-live-trader/pkg/request"
	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/ando9527/poe-live-trader/pkg/ws/pool"
	"github.com/ando9527/poe-live-trader/pkg/ws/server"
)

type Client struct{
	env *env.Client
	database Database
	wsPool types.WsPool
	idCache IDCache
	notifier Notifier
	httpClient types.HttpClient
	wsServer WsServer
	grpcGateway GrpcGateway
}

func NewClient(cfg *env.Client) *Client {
	ctx:=context.Background()
	return &Client{
		env:      cfg,
		database: db.NewClient(),
		wsPool: pool.NewClient(ctx, pool.Config{
			POESSID: cfg.Poesessid,
			League:  cfg.League,
			Filter:  cfg.Filter,
		}),
		idCache:    cache.NewClient(),
		notifier:   notifier.NewClient(ctx),
		httpClient: request.NewClient(),
		wsServer:   server.NewServer("127.0.0.1:8881"),
	}
}

func (c *Client)Run()  {
	c.database.Connect("sqlite.db")
	c.database.Migration()
	c.idCache.Run()
	c.wsPool.Run()
	c.notifier.Run()
	c.wsServer.Run()
	c.grpcGateway.Run()

	for v:=range c.wsPool.GetBuilderChannel(){
		go func(){
			builder, e := v.SetWhisper(c.httpClient)
			if e != nil {
				return
			}
			builder.SetUserID()

			for _,item:=range builder.Build(){
				if allow := c.idCache.AllowSend(item.GetUserID());!allow{
					return
				}

				if ignored:= c.database.IsIgnored(item.GetUserID());ignored{
					return
				}
				c.notifier.SendToQueue(item.GetNotification())
				c.wsServer.SendToServer(item.GetNotification())
			}
		}()
	}

}

