package main

import (
	"log"
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func mainWindow() fyne.Window {
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
			hello.SetText("")
			newSoundWindow()
			hello.SetText("")
		}),
	))

	return w
}

func newSoundWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("New Sound - Go Soundboard")

	hello := widget.NewLabel("Hello, World!")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter new sound name here")
	content := container.NewVBox(input, widget.NewButton("Save", func() {
		log.Println("Content was:", input.Text)
	}))
	w.SetContent(container.NewVBox(
		hello,
		content,
	))

	return w
}
