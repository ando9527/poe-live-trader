package server

import (
	"fmt"

	"github.com/99designs/gqlgen/client"
)

func UpdateSSID(url string, poessid string, user string, pass string)(err error) {
	c := client.New(url)
	query:=fmt.Sprintf(`
	mutation {
		createOrUpdateSSID(input: { Content: "%s" }) {
			Content
		}
	}
`,poessid)
	var resp interface{}
	err = c.Post(query,&resp )
	if err != nil {
		return err
	}
	return nil
}

func GetPOESSID(cloudURL string, user string, pass string) (ssid string, err error) {

	query:=`
	query  {
	  ssid {
		Content
	  }
	}
`
	resp:=struct{
		Ssid struct{
			Content string
		}
	}{}

	c:=client.New(cloudURL)
	err = c.Post(query, &resp)
	if err != nil {
		return "", err
	}

	return resp.Ssid.Content, nil
}

