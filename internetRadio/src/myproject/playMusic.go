package main

import (
	"log"
	"os"
	"time"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func playAudio() {

	f, err := os.Open("./resources/SUBCITY BRB ADAM.wav")

	f2, err := os.Open("./resources/Robyn - Talk To Me.wav")

	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	s2, _, err2 := wav.Decode(f2)

	if err != nil || err2 != nil {
		log.Fatal(err)
	}

	defer streamer.Close()
	defer s2.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	//loop := beep.Loop(-1, streamer)

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

}

func playIdent() {

	f2, err := os.Open("./resources/Robyn - Talk To Me.wav")

	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f2)

	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	//loop := beep.Loop(-1, streamer)

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

}

func main() {
	var y int = 0

	for y != 100 {
		if y%2 == 0 {
			y += 1
			playAudio()
		} else {
			y += 1
			playIdent()
		}
	}
}
