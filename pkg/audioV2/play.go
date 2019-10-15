package audio

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/sirupsen/logrus"
)

type Client struct{
}

func NewClient() *Client {
	c:=&Client{}
	return c
}


func (c *Client)Init()(error error){
	f, err := os.Open(fmt.Sprintf("media/%s.wav", "sample_rate"))
	if err != nil {
		return err
	}

	_, format, err := wav.Decode(f)
	if err != nil {
		return err
	}

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return fmt.Errorf("sound initial failed, %w", err)
	}
	return nil
}
func (c *Client)Play(name string, volume float64) {
	logrus.Debug("Sound playing")
	f, err := os.Open(fmt.Sprintf("media/%s.wav", name))
	if err != nil {
		logrus.Error(err)
	}

	streamer, _, err := wav.Decode(f)
	if err != nil {
		logrus.Error(err)
	}
	defer streamer.Close()


	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer), Paused: false}
	v := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   volume,
		Silent:   false,
	}
	speedy := beep.ResampleRatio(4, 1, v)
	//speaker.Play(speedy)
	done := make(chan bool)
	speaker.Play(beep.Seq(speedy, beep.Callback(func() {
		done <- true
	})))

	select {
		case <-done:
			return
		case <- time.After(time.Second*3):
			logrus.Error("audio time out")
	}
}


