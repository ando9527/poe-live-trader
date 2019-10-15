package requestV2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/stretchr/testify/assert"
)

func expectResult() (itemDetail types.ItemDetail) {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "poe-live-trader") {
		wd = filepath.Dir(wd)
	}
	filePath := fmt.Sprintf("%s/pkg/request/testing/itemDetail.json", wd)
	body, e := ioutil.ReadFile(filePath)
	if e != nil {
		panic(e)
	}
	itemDetail = types.ItemDetail{}
	e = json.Unmarshal(body, &itemDetail)
	if e != nil {
		panic(e)
	}
	return itemDetail
}

func TestClient_RequestItemDetail(t *testing.T) {
	server := NewFakeRequestServer()
	defer server.Close()
	client := NewFakeRequestClient(server.URL)
	actual ,e:= client.RequestItemDetail(types.ItemStub{
		ID:     []string{"6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"},
		Filter: "",
	})
	if e != nil {
		t.Fatal(e)
	}
	expect := expectResult()
	assert.Equal(t, expect, actual)

}

func TestGetPOESSID(t *testing.T) {

}

