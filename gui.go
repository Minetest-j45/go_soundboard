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
			testFile, err := os.Open(file.Text)
			if err != nil {
				hello.SetText("Invalid file")
			} else {
				mainWindowSetContext(fynewindow)
				//confNewSound(name.Text, file.Text)
				defer testFile.Close()
			}
		}),
	))
}

func deleteSoundWindowContext(fynewindow fyne.Window) {

}

func mainWindowSetContext(fynewindow fyne.Window) {
	hello := widget.NewLabel("Hello, World!")
	
	vbox := container.NewVBox()
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWindowSetContext(fynewindow)
			hello.SetText("Hello, World!")
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			newSoundWindowSetContext(fynewindow)
			hello.SetText("Making a new sound!")
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			deleteSoundWindowContext(fynewindow)
			hello.SetText("Deleting a new sound!")
		}),
		widget.NewToolbarAction(theme.MediaRecordIcon(), func() {
			//set record sound window
		}),
		widget.NewToolbarAction(theme./*MediaStopIcon*/VisibilityIcon(), func() {
			//dark mode
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() {
			//light mode
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
	)

	fynewindow.SetContent(container.NewVBox(
		bar,
		vbox,
		hello,
		widget.NewButton("Hi", func() {
			hello.SetText("Recording")
			recordAudio()
			hello.SetText("Playing")
			// playAudio(audio)
		}),
		/*widget.NewButton("+", func() {
			//new sound window
			newSoundWindowSetContext(fynewindow, fyne.CurrentApp())
			hello.SetText("Making a new sound!")
		}),*/
	))
}

func mainWindow(fyneapp fyne.App) fyne.Window {
	w := fyneapp.NewWindow("Go Soundboard - 0.0")

	mainWindowSetContext(w)

	return w
}

func main() {
	a := app.NewWithID("minetest-j45.go_soundboard")
	a.SetIcon(theme.FyneLogo())

	//main window
	w := mainWindow(a)
	//show main window
	w.Resize(fyne.NewSize(700, 400))
	w.ShowAndRun()
}
