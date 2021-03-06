package key

import (
	"context"
	"sync"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/audio"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"github.com/sirupsen/logrus"
)

//func insertByAHK(message string)(err error){
//	logrus.Debug("Auto inserting!")
//	e := clipboard.WriteAll(message)
//	if e != nil {
//		logrus.Error("Copy to clipboard", message)
//	}
//	cmd := exec.Command("./ahk/insert.exe", message)
//	e = cmd.Run()
//	if e != nil {
//		logrus.Error("ahk insert failed, " ,e)
//	}
//
//	audio.Play("audio", -5)
//	return nil
//}

func (c *Client)InsertByRobotGo(message string)(err error){
	logrus.Debug("Auto inserting!")
	e := clipboard.WriteAll(message)
	if e != nil {
		logrus.Error("Copy to clipboard", message)
	}

	title := robotgo.GetTitle()
	if title =="Path of Exile"{
		robotgo.KeyTap("enter")
		robotgo.KeyTap("a",  "control")
		robotgo.KeyTap("v",  "control")
		robotgo.KeyTap("enter")
	}else{
		logrus.Debug("game window is not activated, skip auto pm in game.")
	}


	c.audio.Play("audio", -5)
	return nil
}

type Client struct {
	Message chan string
	ctx context.Context
	Running bool
	sync.Mutex
	audio *audio.Client
}

func NewClient(ctx context.Context)(c *Client){
	c=&Client{
		Message: make(chan string),
		ctx:     ctx,
		Running: true,
		audio:   audio.NewClient(),
	}
	return c
}


func (c *Client) Run(){
	go func(){
		for{
			select {
				case m:=<-c.Message:
					if !c.Running{
						continue
					}
					err := c.InsertByRobotGo(m)
					if err != nil {
						logrus.Error(err)
					}

			case <-c.ctx.Done():
					logrus.Debug("Interrupt keyboard simulator")
					return

			}

		}
	}()

	go func(){
		for{
			keve := robotgo.AddEvent("f2")
			if keve {
				c.Mutex.Lock()
				c.Running = !c.Running
				if c.Running {
					c.audio.Play("on", 0)
				}else{
					c.audio.Play("off", 0)
				}
				c.Mutex.Unlock()

				logrus.Debug("Keyboard simulator is ", c.Running)
				time.Sleep(time.Millisecond*500)
			}

		}
	}()
}