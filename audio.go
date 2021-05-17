package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/go-mp3"
	"github.com/youpy/go-wav"

	"github.com/gen2brain/malgo"
)

func recordAudio() {
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

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1

	var playbackSampleCount uint32
	var capturedSampleCount uint32
	var pCapturedSamples []byte

	f, err := os.OpenFile("samples.hbaj", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		sampleCount := framecount * deviceConfig.Capture.Channels * sizeInBytes
		newCapturedSampleCount := capturedSampleCount + sampleCount
		pCapturedSamples = append(pCapturedSamples, pSample...)
		capturedSampleCount = newCapturedSampleCount

		f.Write(pSample)
	}

	fmt.Println("Recording...")
	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Press Enter to stop recording...")
	fmt.Scanln()

	device.Uninit()

	f, err = os.Open("samples.hbaj")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	onSendFrames := func(pSample, nil []byte, framecount uint32) {
		samplesToRead := framecount * deviceConfig.Playback.Channels * sizeInBytes
		if samplesToRead > capturedSampleCount-playbackSampleCount {
			samplesToRead = capturedSampleCount - playbackSampleCount
		}

		f.Read(pSample)

		playbackSampleCount += samplesToRead
		if playbackSampleCount == uint32(len(pCapturedSamples)) {
			playbackSampleCount = 0
		}
	}

	fmt.Println("Playing...")
	playbackCallbacks := malgo.DeviceCallbacks{
		Data: onSendFrames,
	}

	device, err = malgo.InitDevice(ctx.Context, deviceConfig, playbackCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Press Enter to quit...")
	fmt.Scanln()

	device.Uninit()
}

func playAudio(audioFile string) {
	if audioFile == "" {
		fmt.Println("No audio file selected")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var reader io.Reader
	var channels, sampleRate uint32

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
	defer device.Uninit()

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Press Enter to quit...")
	fmt.Scanln()
}
