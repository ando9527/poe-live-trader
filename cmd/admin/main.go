package main

import (
	"context"

	"github.com/ando9527/poe-live-trader/cmd/admin/conf"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main(){
	e := godotenv.Load("admin.env")
	if e != nil {
		logrus.Fatal(e)
	}
	conf.NewConfig()
	ctx:=context.Background()




	if e != nil {
		logrus.Fatal(e)
	}
	c, err := cloud.NewClient(ctx)
	if err != nil {
		logrus.Fatal(err)
	}
	err = c.UpdateInsert("123")
	if err != nil {
		logrus.Fatal(err)
	}

}


