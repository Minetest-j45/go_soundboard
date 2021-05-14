package main

import (
	"log"
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newSoundWindowSetContext(fynewindow fyne.Window) {
	//new sound window func
	hello := widget.NewLabel("Hello, World!")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter new sound name here")
	fynewindow.SetContent(container.NewVBox(
		hello,
		input,
		widget.NewButton("Cancel", func() {
			mainWindowSetContext(fynewindow)
		}),
		widget.NewButton("Finish", func() {
			log.Println("Content was:", input.Text)
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
            		hello.SetText("Hello, World!")
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
