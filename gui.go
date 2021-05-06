package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func helloWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello, World!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi", func() {
			hello.SetText("Recording")
			/*audio :=*/ recordAudio()
			hello.SetText("Playing")
			// playAudio(audio)
		}),
	))

	return w
}
