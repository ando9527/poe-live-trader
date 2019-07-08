package audio

import (
	"log"
	"os"
	"time"

	"github.com/ando9527/poe-live-trader/conf"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/sirupsen/logrus"
)

func Play() {
	f, err := os.Open("audio.wav")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		logrus.Fatal("sound init failed")
	}
	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   conf.Env.Volume,
		Silent:   false,
	}
	speedy := beep.ResampleRatio(4, 1, volume)
	//speaker.Play(speedy)
	done := make(chan bool)
	speaker.Play(beep.Seq(speedy, beep.Callback(func() {
		done <- true
	})))

	<-done
}
