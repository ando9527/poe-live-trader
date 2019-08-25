package p

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var fakeData = "123"

func insertFakeData(){
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, conf.GoogleProjectId)
	if err != nil {
		logrus.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	err = cloud.UpdateInsert(ctx, client, fakeData)
	if err != nil {
		panic(err)
	}
}

func TestSSID(t *testing.T) {
	insertFakeData()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.SetBasicAuth(conf.User, conf.Pass)
	recorder := httptest.NewRecorder()
	server := http.HandlerFunc(C)
	server.ServeHTTP(recorder,req)
	resp:=recorder.Result()
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}

	assert.Equal(t, fakeData, string(bytes))

}


func TestPOST(t *testing.T) {
	formData := url.Values{}
	formData.Add("poessid", fakeData)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	req.SetBasicAuth(conf.User, conf.Pass)
	recorder := httptest.NewRecorder()
	server := http.HandlerFunc(C)
	server.ServeHTTP(recorder,req)
	resp:=recorder.Result()
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}
	assert.Equal(t, "success", string(bytes))


	ctx := context.Background()
	client, err := firestore.NewClient(ctx, conf.GoogleProjectId)
	if err != nil {
		logrus.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	ssid, err := cloud.QuerySSID(ctx, client)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, fakeData, ssid)

}