package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/malgo"
)

func makeJsonToButtons(fynewindow fyne.Window) []fyne.CanvasObject {
	buttons := openJson()
	var btns []fyne.CanvasObject
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
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			//refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),
	)
	btns = append(btns, bar)
	for _, btn := range buttons.Buttons {
		newbtn := widget.NewButton(btn.Name, func() {
			playAudio(btn.File, fynewindow)
		})
		btns = append(btns, newbtn)
	}
	return btns
}

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
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			//refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter new sound name here")
	file := widget.NewEntry()
	file.SetPlaceHolder("Enter new sound file here")

	cancel := widget.NewButton("Cancel", func() {
		mainWindowSetContext(fynewindow)
	})
	finish := widget.NewButton("Finish", func() {
		log.Println("Name was:", name.Text)
		log.Println("File was:", file.Text)
		testFile, err := os.Open(file.Text)
		if err != nil {
			log.Println("Invalid file")
		} else {
			defer testFile.Close()
		}
		switch strings.ToLower(filepath.Ext(file.Text)) {
		case ".wav":
			mainWindowSetContext(fynewindow)
			confNewSound(name.Text, file.Text)
		case ".mp3":
			mainWindowSetContext(fynewindow)
			confNewSound(name.Text, file.Text)
		case ".hbaj":
			mainWindowSetContext(fynewindow)
			confNewSound(name.Text, file.Text)
		default:
			log.Println("Invalid file extension, we only support .wav, .mp3 and our custom format(.hbaj)")
			return
		}
	})
	fynewindow.SetContent(container.NewVBox(
		bar,
		name,
		file,
		cancel,
		finish,
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
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			//refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter the name of the sound you want to delete here")

	cancel := widget.NewButton("Cancel", func() {
		mainWindowSetContext(fynewindow)
	})

	delete := widget.NewButton("Delete", func() {
		confDeleteSound(name.Text, fynewindow)
		log.Println("Name was:", name.Text)
	})

	fynewindow.SetContent(container.NewVBox(
		bar,
		name,
		cancel,
		delete,
	))
}

func recordSoundWindowContext(fynewindow fyne.Window) {
	log.Println("Record a sound")

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
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			//refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter the name that you want the recording to be saved to here")

	cancel := widget.NewButton("Cancel", func() {
		mainWindowSetContext(fynewindow)
	})

	record := widget.NewButton("Record", func() {
		recordAudio(fynewindow, name.Text)
	})
	fynewindow.SetContent(container.NewVBox(
		bar,
		name,
		cancel,
		record,
	))
}

func recordingSoundWindowContext(fynewindow fyne.Window, device *malgo.Device, saveTo string) {
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
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			//refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),
	)

	stop := widget.NewButton("Finish", func() {
		device.Uninit()
		confNewSound(saveTo, "./recordings/"+saveTo+".hbaj")
		playAudio("./recordings/"+saveTo+".hbaj", fynewindow)
		mainWindowSetContext(fynewindow)

	})
	fynewindow.SetContent(container.NewVBox(
		bar,
		stop,
	))
}

func mainWindowSetContext(fynewindow fyne.Window) {
	log.Println("Main menu")

	buttons := makeJsonToButtons(fynewindow)
	fynewindow.SetContent(container.NewVBox(
		buttons...,
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
