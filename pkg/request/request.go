package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/sirupsen/logrus"
)

type Client struct {
	GetHTTPServerURL func(stub types.ItemStub) (serverURL string)
}

func (client *Client) RequestItemDetail(stub types.ItemStub) (itemDetail types.ItemDetail, err error) {
	logrus.Debug("Requesting data from http url")
	url := client.GetHTTPServerURL(stub)
	resp, err := http.Get(url)
	if err != nil {
		return itemDetail, fmt.Errorf("request url %s failed, %v",url, err)
	}
	if resp == nil || resp.Body == nil {
		return itemDetail,fmt.Errorf("http response is null, url: %s", url)
	}
	defer resp.Body.Close()
	itemDetail = types.ItemDetail{}
	err = json.NewDecoder(resp.Body).Decode(&itemDetail)

	if err != nil {
		return itemDetail, fmt.Errorf("failed to decode json of item detail, %s", url)
	}
	return itemDetail, nil
}

func NewRequestClient() (client *Client) {
	client = &Client{}
	client.GetHTTPServerURL = func(stub types.ItemStub) (serverURL string) {
		return fmt.Sprintf("https://www.pathofexile.com/api/trade/fetch/%s?query=%s", strings.Join(stub.ID, ","), stub.Filter)
	}
	return client
}

