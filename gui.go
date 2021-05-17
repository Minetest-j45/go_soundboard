package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newSoundWindowSetContext(fynewindow fyne.Window) {
	//new sound window func
	hello := widget.NewLabel("Hello, World!")
	name := widget.NewEntry()
	name.SetPlaceHolder("Enter new sound name here")
	file := widget.NewEntry()
	file.SetPlaceHolder("Enter new sound file here")
	fynewindow.SetContent(container.NewVBox(
		hello,
		name,
		file,
		widget.NewButton("Cancel", func() {
			mainWindowSetContext(fynewindow)
		}),
		widget.NewButton("Finish", func() {
			log.Println("Name was:", name.Text)
			log.Println("File was:", file.Text)
			//confNewSound(name.Text, file.Text)
			testFile, err := os.Open(file.Text)
			if err != nil {
				hello.SetText("Invalid file")
			} else {
				mainWindowSetContext(fynewindow)
				defer testFile.Close()
			}
		}),
	))
}

func mainWindowSetContext(fynewindow fyne.Window) {
	hello := widget.NewLabel("Hello, World!")
	fynewindow.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi", func() {
			hello.SetText("Recording")
			recordAudio()
			hello.SetText("Playing")
			// playAudio(audio)
		}),
		widget.NewButton("+", func() {
			//new sound window
			newSoundWindowSetContext(fynewindow)
			hello.SetText("Making a new sound!")
		}),
	))
}

func mainWindow(fyneapp fyne.App) fyne.Window {
	w := fyneapp.NewWindow("Go Soundboard")

	mainWindowSetContext(w)

	return w
}

func main() {
	a := app.New()

	//main window
	w := mainWindow(a)
	//show main window
	w.Resize(fyne.NewSize(700, 400))
	w.ShowAndRun()
}
