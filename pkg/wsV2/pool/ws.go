package pool

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/item"
	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/ando9527/poe-live-trader/pkg/ws/client"
	"github.com/sirupsen/logrus"
)

type Config struct{
	POESSID string
	League  string
	Filter  []string
}


type Client struct {
	ctx context.Context
	pool []*client.Client
	ItemBuilderChan chan types.ItemBuilder
	header http.Header
	cfg Config
}

func (p *Client) GetBuilderChannel() <-chan types.ItemBuilder {
	return p.ItemBuilderChan
}

func NewClient(ctx context.Context, cfg Config) *Client {
	return &Client{
		ctx:          ctx,
		pool:         []*client.Client{},
		ItemBuilderChan : make(chan item.Builder),
		header:       nil,
		cfg:          cfg,
	}
}
func (p *Client)getHeader() (header http.Header) {
	header = getSimChromeCookie()
	logrus.Debug("Using local poessid, ", os.Getenv("CLIENT_POESESSID"))
	cfduid, err := getCFDUID()
	if err != nil {
		logrus.Error(err)
		logrus.Error("cfuid is empty right now")
		cookie := fmt.Sprintf("POESESSID=%s", os.Getenv("CLIENT_POESESSID"))
		header.Add("Cookie", cookie)
		return header
	}
	cookie := fmt.Sprintf("__cfduid=%s; POESESSID=%s", cfduid, os.Getenv("CLIENT_POESESSID"))
	header.Add("Cookie", cookie)
	return header
}

func (p *Client)Run(){
	p.header = p.getHeader()
	logrus.Debugf("Using Header: %s", p.header)
	for _,id:=range p.cfg.Filter{
		cfg:=client.Config{
			POESSID: p.cfg.POESSID,
			League:  p.cfg.League,
			Filter:  id,
			Header:  p.header,
		}
		c:=client.NewClient(p.ctx, cfg)
		c.Run()
		p.pool= append(p.pool, c)
		p.merge(c.ItemStub)
	}

}


func (p *Client)merge(cs <-chan types.ItemStub)  {
	go func(){
		for v:=range cs {
			p.ItemStubChan<-v
		}
	}()
}

func getSimChromeCookie() (header http.Header) {
	header = make(http.Header)
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7,zh-CN;q=0.6,ja;q=0.5")
	header.Add("Cache-Control", "no-cache")
	//header.Add("Connection", "Upgrade")
	header.Add("Host", "www.pathofexile.com")
	header.Add("Origin", "https://www.pathofexile.com")
	header.Add("Pragma", "no-cache")
	//header.Add("Sec-WebSocket-Extensions", "permessage-deflate; client_max_window_bits")
	//header.Add("Sec-WebSocket-Key", "Oa+B/nEJMeezec/bNsjTwg==")
	//header.Add("Sec-WebSocket-Version", "13")
	//header.Add("Upgrade", "websocket")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	return header
}

func getCFDUID()(s string, err error){
	c:=http.Client{
		Timeout:       time.Second*3,
	}
	url:="https://www.pathofexile.com"
	resp, err := c.Get(url)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	cookie:=resp.Cookies()
	for _,v:=range cookie{
		if v.Name=="__cfduid"{
			return v.Value, nil
		}
	}
	return "", fmt.Errorf("can't find cfduid in url, %s", url)
}