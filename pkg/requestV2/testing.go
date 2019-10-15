package requestV2

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func NewFakeRequestClient(serverURL string) (client *Client) {
	client = &Client{
		httpClient:       &http.Client{
			Timeout:       time.Second*10,
		},
	}
	client.GetHTTPServerURL = func(idList []string, filterID string) (serverURL2 string) {
		return serverURL
	}
	return client
}

func NewFakeRequestServer() *httptest.Server {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "poe-live-trader") {
		wd = filepath.Dir(wd)
	}
	filePath := fmt.Sprintf("%s/pkg/request/testing/itemDetail.json", wd)
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
