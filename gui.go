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
	log.Println("Making a new sound")

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			newSoundWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			deleteSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.MediaRecordIcon(), func() {
			recordSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.VisibilityIcon(), func() {
			//dark mode
			log.Println("Dark mode")
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() {
			//light mode
			log.Println("Light mode")
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter new sound name here")
	file := widget.NewEntry()
	file.SetPlaceHolder("Enter new sound file here")
	fynewindow.SetContent(container.NewVBox(
		bar,
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
				log.Println("Invalid file")
			} else {
				mainWindowSetContext(fynewindow)
				//confNewSound(name.Text, file.Text)
				defer testFile.Close()
			}
		}),
	))
}

func deleteSoundWindowContext(fynewindow fyne.Window) {
	log.Println("Deleting a sound")

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			newSoundWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			deleteSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.MediaRecordIcon(), func() {
			recordSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.VisibilityIcon(), func() {
			//dark mode
			log.Println("Dark mode")
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() {
			//light mode
			log.Println("Light mode")
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
	)
	
	fynewindow.SetContent(container.NewVBox(
		bar,
	))
}

func recordSoundWindowContext(fynewindow fyne.Window) {
	log.Println("Recording a sound")

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			newSoundWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			deleteSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.MediaRecordIcon(), func() {
			recordSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.VisibilityIcon(), func() {
			//dark mode
			log.Println("Dark mode")
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() {
			//light mode
			log.Println("Light mode")
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
	)
	
	fynewindow.SetContent(container.NewVBox(
		bar,
		widget.NewButton("record", func() {
			recordAudio()
		}),
	))
}

func mainWindowSetContext(fynewindow fyne.Window) {
	log.Println("Main menu")

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			newSoundWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			deleteSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.MediaRecordIcon(), func() {
			recordSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.VisibilityIcon(), func() {
			//dark mode
			log.Println("Dark mode")
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() {
			//light mode
			log.Println("Light mode")
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
	)

	fynewindow.SetContent(container.NewVBox(
		bar,
		widget.NewButton("Hi", func() {
			recordAudio()
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
	a.SetIcon(resourceCUsersUserGosoundboardGosoundboardlogoPng)

	//main window
	w := mainWindow(a)
	//show main window
	w.Resize(fyne.NewSize(700, 400))
	w.ShowAndRun()
}
