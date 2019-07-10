package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
)

func NewFakeRequestClient(serverURL string) (client *Client) {
	client = &Client{}
	client.GetHTTPServerURL = func(itemID []string) (serverURL2 string) {
		return serverURL
	}
	return client
}

func NewFakeRequestServer() *httptest.Server {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "poe-live-trader") {
		wd = filepath.Dir(wd)
	}
	filePath := fmt.Sprintf("%s/pkg/v2/request/testing/itemDetail.json", wd)
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		panic(e)
	}
	f := func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, string(bytes))
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	return server
}
