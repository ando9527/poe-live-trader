package cloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ando9527/poe-live-trader/cmd/client/conf"
	"github.com/sirupsen/logrus"
)

func PostSSID(url string, poessid string, user string, pass string) {
	logrus.Debug("post ", poessid, " to ", url)
	s:=SSID{}
	s.Content = poessid
	ytes, e := json.Marshal(&s)
	if e != nil {
		panic(e)
	}

	request, e := http.NewRequest("POST", url, bytes.NewReader(ytes))
	if e != nil {
		panic(e)
	}
	request.SetBasicAuth(user, pass)
	request.Header.Set("Content-Type", "application/json")
	c:=http.Client{}
	resp, e := c.Do(request)
	if e != nil {
		panic(e)
	}
	if resp.StatusCode ==200{
		fmt.Println("success")
	}
}

func GetPOESSID(url string) (ssid string) {

	c := http.Client{Timeout: time.Second * 10}
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		logrus.Panic(e)
	}
	req.SetBasicAuth(conf.Env.User, conf.Env.Pass)
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	s := SSID{}


	er := json.NewDecoder(resp.Body).Decode(&s)
	defer resp.Body.Close()
	if er != nil {
		logrus.Panic(er)
	}
	logrus.Debug("using cloud poessid")
	return s.Content
}

