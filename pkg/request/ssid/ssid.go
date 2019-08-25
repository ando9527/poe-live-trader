package ssid

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct{
	*http.Client
}

func (c*Client) GetSSID()(ssid string, err error) {
	url:="http://127.0.0.1:8000"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	request.SetBasicAuth(os.Getenv("APP_USER"), os.Getenv("APP_PASS"))
	resp, err := c.Do(request)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	ssid = string(bytes)
	return ssid, nil
}

func NewClient()(c *Client){
	c = &Client{&http.Client{
		Timeout:       time.Second*10,
	}}
	return c
}