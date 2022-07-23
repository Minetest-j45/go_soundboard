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
)

func makeJsonToButtons(fynewindow fyne.Window) []fyne.CanvasObject {
	buttons := openJson()
	var btns []fyne.CanvasObject
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() { //home
			mainWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { //add
			newSoundWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() { //remove
			deleteSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.VisibilityIcon(), func() { //dark mode
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() { //light mode
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
		/*widget.NewToolbarAction(theme.ViewRefreshIcon(), func() { //refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),*/
	)
	btns = append(btns, bar)
	for _, btn := range buttons.Buttons {
		file := btn.File
		newbtn := widget.NewButton(btn.Name, func() {
			go playAudio(file, fynewindow)
		})
		btns = append(btns, newbtn)
	}
	return btns
}

func newSoundWindowSetContext(fynewindow fyne.Window) {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() { //home
			mainWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { //add
			newSoundWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() { //remove
			deleteSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.VisibilityIcon(), func() { //dark mode
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() { //light mode
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
		/*widget.NewToolbarAction(theme.ViewRefreshIcon(), func() { //refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),*/
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

		exists := confExists(name.Text)
		if exists {
			log.Println("A button with the name `" + name.Text + "` already exists")
			return
		}

		testFile, err := os.Open(file.Text)
		if err != nil {
			log.Println(err)
			return
		}
		defer testFile.Close()

		switch strings.ToLower(filepath.Ext(file.Text)) {
		case ".wav":
			mainWindowSetContext(fynewindow)
			confNewSound(name.Text, file.Text)
			mainWindowSetContext(fynewindow)
		case ".mp3":
			mainWindowSetContext(fynewindow)
			confNewSound(name.Text, file.Text)
			mainWindowSetContext(fynewindow)
		default:
			log.Println("Invalid file extension, we only support .wav and .mp3 ")
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
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() { //home
			mainWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { //add
			newSoundWindowSetContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() { //remove
			deleteSoundWindowContext(fynewindow)
		}),
		widget.NewToolbarAction(theme.VisibilityIcon(), func() { //dark mode
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewToolbarAction(theme.VisibilityOffIcon(), func() { //light mode
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		}),
		/*widget.NewToolbarAction(theme.ViewRefreshIcon(), func() { //refresh
			log.Println("Refresh")
			mainWindowSetContext(fynewindow)
		}),*/
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

func mainWindowSetContext(fynewindow fyne.Window) {
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
	a := app.NewWithID("com.minetest-j45.go_soundboard")
	a.SetIcon(resourceCUsersUserGosoundboardGosoundboardlogoPng)

	w := mainWindow(a)
	w.Resize(fyne.NewSize(700, 400))
	w.ShowAndRun()
}
