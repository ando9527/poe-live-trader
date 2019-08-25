package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var fakeData = "123"

func insertFakeData(){
	ctx:=context.Background()
	c, err := cloud.NewClient(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	err = c.UpdateInsert( fakeData)
	if err != nil {
		panic(err)
	}
}

func TestSSID(t *testing.T) {
	insertFakeData()
	srv:=Server{}
	srv.router = http.NewServeMux()
	srv.Routes()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.SetBasicAuth(conf.User, conf.Pass)
	recorder := httptest.NewRecorder()
	srv.router.ServeHTTP(recorder, req)
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
	srv:=Server{}
	srv.router = http.NewServeMux()
	srv.Routes()
	srv.router.ServeHTTP(recorder,req)
	resp:=recorder.Result()
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}
	assert.Equal(t, "success", string(bytes))


	ctx := context.Background()
	client, err := cloud.NewClient(ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.Close()
	ssid, err := client.QuerySSID()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, fakeData, ssid)

}