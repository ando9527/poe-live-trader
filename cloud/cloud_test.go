package p

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)



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
	assert.Equal(t, "123", string(bytes))

}