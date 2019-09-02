package types

import (
	"time"
)

const ANCHOR = "ANCHOR"

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

type ItemID struct {
	New []string `json:"new"`
}
