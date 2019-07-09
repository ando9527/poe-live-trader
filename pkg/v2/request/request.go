package request

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ando9527/poe-live-trader/pkg/v2/types"
)

type Client struct {
}

func (client *Client) RequestItemDetail(strings []string) (itemDetail types.ItemDetail) {
	body, e := ioutil.ReadFile("testing/itemDetail.json")
	if e != nil {
		panic(e)
	}
	itemDetail = types.ItemDetail{}
	e = json.Unmarshal(body, &itemDetail)
	if e != nil {
		panic(e)
	}
	return itemDetail
}

func NewRequestClient() (client *Client) {
	return &Client{}
}
