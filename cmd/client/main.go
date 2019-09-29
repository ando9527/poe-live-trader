package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/ando9527/poe-live-trader/cmd/client/env"
	"github.com/ando9527/poe-live-trader/pkg/audio"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/ando9527/poe-live-trader/pkg/trader"
	"github.com/ando9527/poe-live-trader/pkg/ws"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	//_ "net/http/pprof"
)

var (
	version  string
)

func pause() {
	fmt.Print("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func NotifyDC(cancel context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			select {
			case <-interrupt:
				cancel()
				time.Sleep(time.Second*2)
				return
			}
		}
	}()

}

func main() {
	//go func() {
	//	logrus.Info(http.ListenAndServe("localhost:6060", nil))
	//}()
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
		League:      cfg.League,
		Filter:      cfg.Filter,
	}

	ctx, cancel := context.WithCancel(context.Background())
	NotifyDC(cancel)
	client := trader.NewTrader(ctx, config)

	client.Launch()


	whisper := client.Whisper

	for {
		select {
		case result := <-whisper:
			logrus.Info(result)
			client.Mutex.Lock()
			if client.IDCache[getName(result)]{
				logrus.Debug("duplicated user in cache, ",getName(result))
				continue
			}
			client.KeySim.Message<-result
			client.IDCache[getName(result)]=true

			client.Mutex.Unlock()
			audio.Play("audio", cfg.Volume)
			//if err != nil {
			//	logrus.Warn("failed copy whisper to clipboard.")
			//}
		}
	}

}

func getName(template string)(n string){
	return strings.Split(template, " ")[0]
}