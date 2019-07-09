package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/audio"
	"github.com/atotto/clipboard"

	"github.com/ando9527/poe-live-trader/conf"
	"github.com/sirupsen/logrus"
)

type ItemDetail struct {
	Result []struct {
		ID      string `json:"id"`
		Listing struct {
			Method  string    `json:"method"`
			Indexed time.Time `json:"indexed"`
			Stash   struct {
				Name string `json:"name"`
				X    int    `json:"x"`
				Y    int    `json:"y"`
			} `json:"stash"`
			Whisper string `json:"whisper"`
			Account struct {
				Name              string `json:"name"`
				LastCharacterName string `json:"lastCharacterName"`
				Online            struct {
					League string `json:"league"`
				} `json:"online"`
				Language string `json:"language"`
			} `json:"account"`
			Price struct {
				Type     string `json:"type"`
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			} `json:"price"`
		} `json:"listing"`
		Item struct {
			Verified   bool   `json:"verified"`
			W          int    `json:"w"`
			H          int    `json:"h"`
			Ilvl       int    `json:"ilvl"`
			Icon       string `json:"icon"`
			League     string `json:"league"`
			Name       string `json:"name"`
			TypeLine   string `json:"typeLine"`
			Identified bool   `json:"identified"`
			Note       string `json:"note"`
			Properties []struct {
				Name        string          `json:"name"`
				Values      [][]interface{} `json:"values"`
				DisplayMode int             `json:"displayMode"`
			} `json:"properties"`
			ExplicitMods []string `json:"explicitMods"`
			DescrText    string   `json:"descrText"`
			FrameType    int      `json:"frameType"`
			StackSize    int      `json:"stackSize"`
			MaxStackSize int      `json:"maxStackSize"`
			Category     struct {
				Currency []interface{} `json:"currency"`
			} `json:"category"`
			Extended struct {
				Text string `json:"text"`
			} `json:"extended"`
		} `json:"item"`
	} `json:"result"`
}

type Handler struct {
}

func (h *Handler) ConvertJSON(message string) (liveData LiveData) {
	liveData = LiveData{}
	err := json.Unmarshal([]byte(message), &liveData)
	if err != nil {
		logrus.Fatalf("decode json message from ws server failed, message: %s", message)
	}
	return liveData
}

func GetHTTPServerURL(itemID []string) (serverURL string) {
	return fmt.Sprintf("https://www.pathofexile.com/api/trade/fetch/%s?query=%s", strings.Join(itemID, ","), conf.Env.Filter)
}
func (h *Handler) GetItemDetail(url string) (itemDetail ItemDetail) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Fatalf("Get item detail from url failed, url: %s", url)
	}
	if resp == nil || resp.Body == nil {
		log.Fatalf("http response is nil, url: %s", url)
	}
	defer resp.Body.Close()
	itemDetail = ItemDetail{}
	err = json.NewDecoder(resp.Body).Decode(&itemDetail)

	if err != nil {
		logrus.Fatalf("failed to decode json of item detail, url: %s", url)
	}
	return itemDetail
}

func (h *Handler) Do(message string) {
	go func() {
		// convert message to json
		liveData := h.ConvertJSON(message)

		// get item detail through http request
		url := GetHTTPServerURL(liveData.New)
		itemDetail := h.GetItemDetail(url)

		for _, result := range itemDetail.Result {
			fmt.Println(result.Listing.Whisper)
			// copy to clipboard
			err := clipboard.WriteAll(result.Listing.Whisper)
			if err != nil {
				logrus.Warn("failed copy whisper to clipboard.")
			}
			// sound alert
			audio.Play()
		}
	}()
}
