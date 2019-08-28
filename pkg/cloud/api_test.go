package cloud

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_handleSSID(t *testing.T) {
	ssid:=SSID{
		Content: fakeData,
	}
	b, e := json.Marshal(ssid)
	if e != nil {
		panic(e)
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b) )
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	f:= server.handleSSID()
	f.ServeHTTP(recorder,req)

	resp:=recorder.Result()
	b2, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	assert.Equal(t, SUCCESS, string(b2))

}