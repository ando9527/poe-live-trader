package audio

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/sirupsen/logrus"
)

func Play(name string, volume float64) {
	logrus.Debug("Sound playing")
	f, err := os.Open(fmt.Sprintf("%s.wav", name))
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
		logrus.Panic("sound init failed")
	}
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

	<-done
}


