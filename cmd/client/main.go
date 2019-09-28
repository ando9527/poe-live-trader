package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ando9527/poe-live-trader/cmd/client/env"
	"github.com/ando9527/poe-live-trader/pkg/audio"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/ando9527/poe-live-trader/pkg/trader"
	"github.com/ando9527/poe-live-trader/pkg/ws"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	version  string
)

func pause() {
	fmt.Print("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	err := godotenv.Load("client.env")
	if err != nil {
		logrus.Error("Error loading client.env file")
		pause()
		return
	}
	cfg, err:= env.NewEnv()
	if err != nil {
		logrus.Error(errors.Wrap(err, "Please setup .env file properly"))
		pause()
		return
	}


	log.InitLogger(cfg.LogLevel)
	logrus.Infof("Poe Live Trader %s", version)

	config:=		ws.Config{
		CloudEnable: cfg.CloudEnable,
		CloudURL:    cfg.CloudUrl,
		User:        cfg.User,
		Pass:        cfg.Pass,
		League:      cfg.League,
		Filter:      cfg.Filter,
	}
	client := trader.NewTrader(config)
	client.Launch()


	whisper := client.Whisper

	for {
		select {
		case result := <-whisper:
			fmt.Println(result)
			audio.Play(cfg.Volume)
			client.Mutex.Lock()
			if client.IDCache[getName(result)]{
				continue
			}
			client.LocalServer.Message<-result
			client.IDCache[getName(result)]=true
			client.Mutex.Unlock()
			//if err != nil {
			//	logrus.Warn("failed copy whisper to clipboard.")
			//}
		}
	}

}

func getName(template string)(n string){
	return strings.Split(template, " ")[0]
}