package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)
var (poessid string
     migration bool
)
func init(){
	flag.BoolVar(&migration, "m", false, "execute migration?")
	flag.StringVar(&poessid, "i", "", "Insert POESSID")
}
func main(){
	e := godotenv.Load("admin.env")
	if e != nil {
		logrus.Panic(e)
	}
	flag.Parse()
	if poessid!=""{
		s:=cloud.SSID{}
		s.Content = poessid
		ytes, e := json.Marshal(&s)
		if e != nil {
			panic(e)
		}
		request, e := http.NewRequest("POST", os.Getenv("APP_CLOUD_URL"), bytes.NewReader(ytes))
		if e != nil {
			panic(e)
		}
		c:=http.Client{}
		resp, e := c.Do(request)
		if e != nil {
			panic(e)
		}
		if resp.StatusCode ==200{
			fmt.Println("success")
		} 
	}

	if migration==true{
		s := cloud.NewServer()
		s.Connect()
		s.Migration()
	}
	//

	//conf.NewConfig()
	//ctx:=context.Background()
	//
	//
	//if e != nil {
	//	logrus.Panic(e)
	//}
	//c, err := cloud.NewClient(ctx)
	//if err != nil {
	//	logrus.Panic(err)
	//}
	//err = c.UpdateInsert(poessid)
	//if err != nil {
	//	logrus.Panic(err)
	//}
	//fmt.Println("success!")

}


