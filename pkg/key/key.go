package key

import (
	"context"
	"os/exec"
	"sync"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/audio"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"github.com/sirupsen/logrus"
)
type Client struct {
	Message chan string
	ctx context.Context
	Running bool
	sync.Mutex
}

func NewClient(ctx context.Context)(c *Client){
	c=&Client{
		Message: make(chan string),
		ctx:     ctx,
		Running: true,
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
					logrus.Debug("Auto inserting!")
					cmd := exec.Command("./ahk/insert.exe", m)
					e := cmd.Run()
					if e != nil {
						logrus.Error("ahk insert failed, " ,e)
					}
					e = clipboard.WriteAll(m)
					if e != nil {
						logrus.Error("Copy to clipboard", m)
					}
					audio.Play("audio", -5)

					//robotgo.WriteAll(m)
					//robotgo.KeyTap("enter")
					//robotgo.KeyTap("a","ctrl")
					//robotgo.KeyTap("v","ctrl")
					//robotgo.KeyTap("enter")
				case <-c.ctx.Done():
					logrus.Debug("Interrupt keyboard simulator")
					return

			}

		}
	}()

	go func(){
		for{
			c.Mutex.Lock()
			keve := robotgo.AddEvent("f2")
			if keve {
				c.Running = !c.Running

				if c.Running {
					audio.Play("on", 0)
				}else{
					audio.Play("off", 0)
				}

				logrus.Debug("Keyboard simulator is ", c.Running)
				time.Sleep(time.Millisecond*500)
			}
			c.Mutex.Unlock()
		}
	}()
}