package main

import (
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"

	"github.com/gen2brain/malgo"
)

func playAudio(audioFile string, fynewindow fyne.Window) {
	if audioFile == "" {
		log.Println("No audio file selected")
		return
	}

	file, err := os.Open(audioFile)
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()

	var reader io.Reader
	var channels, sampleRate uint32
	var dur time.Duration

	switch strings.ToLower(filepath.Ext(audioFile)) {
	case ".wav":
		w := wav.NewReader(file)
		f, err := w.Format()
		if err != nil {
			log.Println(err)
			return
		}

		reader = w
		channels = uint32(f.NumChannels)
		sampleRate = f.SampleRate

		dur, err = w.Duration()
		if err != nil {
			log.Println(err)
			return
		}
	case ".mp3":
		m, err := mp3.NewDecoder(file)
		if err != nil {
			log.Println(err)
			return
		}

		reader = m
		channels = 2
		sampleRate = uint32(m.SampleRate())

		sampleSize := 4                              // From documentation.
		samples := int(m.Length()) / sampleSize      // Number of samples.
		audioLength := int(samples) / m.SampleRate() // Audio length in seconds.
		dur = time.Duration(audioLength * int(math.Pow(10, 9)))
	default:
		log.Println("Not a valid file.")
		return
	}

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		log.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = channels
	deviceConfig.SampleRate = sampleRate
	deviceConfig.Alsa.NoMMap = 1

	var device *malgo.Device

	// This is the function that's used for sending more data to the device for playback.
	onSamples := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		io.ReadFull(reader, pOutputSample)
	}

	deviceCallbacks := malgo.DeviceCallbacks{
		Data: onSamples,
	}
	device, err = malgo.InitDevice(ctx.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		log.Println(err)
		return
	}

	err = device.Start()
	if err != nil {
		log.Println(err)
		return
	}

	select {
	case <-time.After(dur):
		device.Stop()
		device.Uninit()
	}
}
