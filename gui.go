package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func helloWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("Go Soundboard")

	hello := widget.NewLabel("Hello, World!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi", func() {
			hello.SetText("Recording")
			/*audio :=*/ recordAudio()
			hello.SetText("Playing")
			// playAudio(audio)
		}),
		widget.NewButton("+", func() {
			//hello.SetText("Recording")
			//make new sound/button newSound()
			//hello.SetText("Playing")
			// playAudio(audio)
		}),
	))

	return w
}
