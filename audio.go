package main

import (
	"io"
	"os"
	"os/signal"

	"github.com/gordonklaus/portaudio"
)

func recordAudio() []int32 {
	var out []int32

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		panic(err)
	}
	defer stream.Stop()

	for {
		if err := stream.Read(); err != nil {
			panic(err)
		}

		out = append(out, in...)
		select {
		case <-sig:
			return out
		default:
		}
	}
}

func playAudio(data []int32) {
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(data), &data)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		panic(err)
	}
	defer stream.Stop()

	if err == io.EOF {
		return
	}
	if err != nil {
		panic(err)
	}

	stream.Write()
}
