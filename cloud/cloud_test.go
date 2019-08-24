package p

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"cloud.google.com/go/firestore"
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

	err = updateInsert(ctx, client, fakeData)
	if err != nil {
		panic(err)
	}
}

func TestSSID(t *testing.T) {
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