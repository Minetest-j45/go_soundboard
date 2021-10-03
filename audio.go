package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"

	"github.com/gen2brain/malgo"
)

func playAudio(audioFile string, fynewindow fyne.Window) {
	if audioFile == "" {
		fmt.Println("No audio file selected")
		os.Exit(1)
	}

	file, err := os.Open(audioFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	switch strings.ToLower(filepath.Ext(audioFile)) {
	case ".wav":
		playNormalAudio(audioFile, fynewindow)

	case ".mp3":
		playNormalAudio(audioFile, fynewindow)
	default:
		fmt.Println("Not a valid file.")
		os.Exit(1)
	}
}

func playNormalAudio(audioFile string, fynewindow fyne.Window) {
	if audioFile == "" {
		fmt.Println("No audio file selected")
		os.Exit(1)
	}

	file, err := os.Open(audioFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var reader io.Reader
	var channels, sampleRate uint32
	//log.Println(strings.ToLower(filepath.Ext(audioFile)))
	switch strings.ToLower(filepath.Ext(audioFile)) {
	case ".wav":
		w := wav.NewReader(file)
		f, err := w.Format()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader = w
		channels = uint32(f.NumChannels)
		sampleRate = f.SampleRate

	case ".mp3":
		m, err := mp3.NewDecoder(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader = m
		channels = 2
		sampleRate = uint32(m.SampleRate())
	default:
		fmt.Println("Not a valid file.")
		os.Exit(1)
	}

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

	// This is the function that's used for sending more data to the device for playback.
	onSamples := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		io.ReadFull(reader, pOutputSample)
	}

	deviceCallbacks := malgo.DeviceCallbacks{
		Data: onSamples,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println("Playing a sound")
	stop := widget.NewButton("Stop", func() {
		device.Uninit()
		mainWindowSetContext(fynewindow)
	})
	fynewindow.SetContent(container.NewVBox(
		stop,
	))
}
