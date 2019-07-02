package audio

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"

	"github.com/faiface/beep/speaker"
)

var streamer beep.StreamSeekCloser

func init() {

}

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

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
