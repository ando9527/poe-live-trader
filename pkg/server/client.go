package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/machinebox/graphql"
)
func basicAuth(header *http.Header, username, password string)  {
	auth := username + ":" + password
	header.Add("Authorization", fmt.Sprintf("Basic %s",base64.StdEncoding.EncodeToString([]byte(auth))))
}

func UpdateSSID(url string, poessid string, user string, pass string)(err error) {
	client := graphql.NewClient(url)

	req := graphql.NewRequest(`
	mutation ($key: String!){
	  createOrUpdateSSID(input: { Content: $key }) {
		Content
	  }
	}
`)
	req.Var("key", poessid)
	basicAuth(&req.Header, user,pass)
	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()
	if err := client.Run(ctx, req, nil); err != nil {
		return err
	}
	return err
}

func GetPOESSID(cloudURL string, user string, pass string) (ssid string, err error) {
	client := graphql.NewClient(cloudURL)
	req := graphql.NewRequest(`
	query  {
	  ssid {
		Content
	  }
	}
`)
	req.Header.Set("Cache-Control", "no-cache")
	basicAuth(&req.Header, user,pass)
	ctx := context.Background()
	resp:=struct{
		Ssid struct{
			Content string
		}
	}{}
	if err := client.Run(ctx, req, &resp); err != nil {
		return "",err
	}
	return resp.Ssid.Content,err
}

