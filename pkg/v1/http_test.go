package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeItemHandler(w http.ResponseWriter, _ *http.Request) {
	bytes, e := ioutil.ReadFile("testing/itemDetail.json")
	if e != nil {
		panic(e)
	}
	_, _ = fmt.Fprint(w, string(bytes))
}

func TestHandler_GetItemDetail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fakeItemHandler))
	defer server.Close()
	ih := Handler{}
	itemDetail := ih.GetItemDetail(server.URL)
	actual := itemDetail.Result[0].ID
	expect := "6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"
	assert.Equal(t, expect, actual)
}
