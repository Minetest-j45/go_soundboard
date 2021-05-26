package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)


func newSoundWindowSetContext(fynewindow fyne.Window, fyneapp fyne.App) {
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
			mainWindowSetContext(fynewindow, fyneapp)
		}),
		widget.NewButton("Finish", func() {
			log.Println("Name was:", name.Text)
			log.Println("File was:", file.Text)
			//confNewSound(name.Text, file.Text)
			testFile, err := os.Open(file.Text)
			if err != nil {
				hello.SetText("Invalid file")
			} else {
				mainWindowSetContext(fynewindow, fyneapp)
				defer testFile.Close()
			}
		}),
	))
}

func deleteSoundWindowContext(fynewindow fyne.Window, fyneapp fyne.App) {

}

func mainWindowSetContext(fynewindow fyne.Window, fyneapp fyne.App) {
	hello := widget.NewLabel("Hello, World!")

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWindowSetContext(fynewindow, fyneapp)
			hello.SetText("Hello, World!")
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			newSoundWindowSetContext(fynewindow, fyneapp)
			hello.SetText("Making a new sound!")
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			deleteSoundWindowContext(fynewindow, fyneapp)
			hello.SetText("Deleting a new sound!")
		}),
		/*widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			//switch between light and dark mode
			fyneapp.Settings().SetTheme(theme.DarkTheme())
			log.Println(fyne.CurrentApp().Settings().Theme())
			fyneapp.Settings().SetTheme(theme.LightTheme())
			log.Println(fyne.CurrentApp().Settings().Theme())
		}),*/
	)

	fynewindow.SetContent(container.NewVBox(
		bar,
		hello,
		widget.NewButton("Hi", func() {
			hello.SetText("Recording")
			recordAudio()
			hello.SetText("Playing")
			// playAudio(audio)
		}),
		/*widget.NewButton("+", func() {
			//new sound window
			newSoundWindowSetContext(fynewindow, fyneapp)
			hello.SetText("Making a new sound!")
		}),*/
	))
}

func mainWindow(fyneapp fyne.App) fyne.Window {
	w := fyneapp.NewWindow("Go Soundboard")

	mainWindowSetContext(w, fyneapp)

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
