package ws

import (
	"context"

	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/ando9527/poe-live-trader/pkg/ws/client"
)

type Config struct{
	POESSID string
	League  string
	Filter  []string
}


type Pool struct {
	ctx context.Context
	pool []*client.Client
	ItemStubChan chan types.ItemStub
	cfg Config
}

func NewPool(ctx context.Context, cfg Config) *Pool {
	return &Pool{
		ctx:          ctx,
		pool:         []*client.Client{},
		ItemStubChan: make(chan types.ItemStub),
		cfg:          cfg,
	}
}

func (p *Pool)Run()(err error){
	for _,id:=range p.cfg.Filter{
		cfg:=client.Config{
			POESSID: p.cfg.POESSID,
			League:  p.cfg.League,
			Filter:  id,
		}
		c:=client.NewClient(p.ctx, cfg)
		err := c.Run()
		if err != nil {
			return err
		}
		p.pool= append(p.pool, c)
		p.merge(c.ItemStub)
	}



	return nil
}


func (p *Pool)merge(cs <-chan types.ItemStub)  {
	go func(){
		for v:=range cs {
			p.ItemStubChan<-v
		}
	}()
}

