package request

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/sirupsen/logrus"
)

type Client struct {
	GetHTTPServerURL func(itemID []string) (serverURL string)
}

func (client *Client) RequestItemDetail(itemID []string) (itemDetail types.ItemDetail) {
	url := client.GetHTTPServerURL(itemID)
	resp, err := http.Get(url)
	if err != nil {
		logrus.Panicf("Get item detail from url failed, url: %s", url)
	}
	if resp == nil || resp.Body == nil {
		log.Fatalf("http response is nil, url: %s", url)
	}
	defer resp.Body.Close()
	itemDetail = types.ItemDetail{}
	err = json.NewDecoder(resp.Body).Decode(&itemDetail)

	if err != nil {
		logrus.Panicf("failed to decode json of item detail, url: %s", url)
	}
	return itemDetail
}

func NewRequestClient(filter string) (client *Client) {
	client = &Client{}
	client.GetHTTPServerURL = func(itemID []string) (serverURL string) {
		return fmt.Sprintf("https://www.pathofexile.com/api/trade/fetch/%s?query=%s", strings.Join(itemID, ","), filter)
	}
	return client
}

