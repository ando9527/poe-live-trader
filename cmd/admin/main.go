package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ando9527/poe-live-trader/cmd/admin/conf"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)
var poessid string
func init(){
	flag.StringVar(&poessid, "p", "", "Insert POESSID")
}
func main(){
	flag.Parse()
	if poessid==""{
		flag.Usage()
	}

	e := godotenv.Load("admin.env")
	if e != nil {
		logrus.Panic(e)
	}
	conf.NewConfig()
	ctx:=context.Background()


	if e != nil {
		logrus.Panic(e)
	}
	c, err := cloud.NewClient(ctx)
	if err != nil {
		logrus.Panic(err)
	}
	err = c.UpdateInsert(poessid)
	if err != nil {
		logrus.Panic(err)
	}
	fmt.Println("success!")

}


