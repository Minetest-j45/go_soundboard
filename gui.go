package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func jsonToBtns(fynewindow fyne.Window) []fyne.CanvasObject {
	buttons := openJson()

	var btns []fyne.CanvasObject

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin(fynewindow)
		}),
	)

	settings := openSettings()
	lay := layout.NewGridLayoutWithColumns(settings.Columns)
	var soundbtns []fyne.CanvasObject

	btns = append(btns, bar)

	for _, btn := range buttons.Buttons {
		file := btn.File
		newbtn := widget.NewButton(btn.Name, func() {
			go playAudio(file, fynewindow)
		})
		soundbtns = append(soundbtns, newbtn)
	}

	grid := container.New(lay, soundbtns...)
	btns = append(btns, grid)

	return btns
}

func addWin(fynewindow fyne.Window) {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin(fynewindow)
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter new sound name here")
	file := widget.NewEntry()
	file.SetPlaceHolder("Enter new sound file here")

	cancel := widget.NewButton("Cancel", func() {
		mainWin(fynewindow)
	})
	finish := widget.NewButton("Finish", func() {
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
			mainWin(fynewindow)
			confNewSound(name.Text, file.Text)
			mainWin(fynewindow)
		case ".mp3":
			mainWin(fynewindow)
			confNewSound(name.Text, file.Text)
			mainWin(fynewindow)
		default:
			log.Println("Invalid file extension, we only support .wav and .mp3")
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

func rmWin(fynewindow fyne.Window) {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin(fynewindow)
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter the name of the sound you want to delete here")

	cancel := widget.NewButton("Cancel", func() {
		mainWin(fynewindow)
	})

	delete := widget.NewButton("Delete", func() {
		confDeleteSound(name.Text, fynewindow)
	})

	fynewindow.SetContent(container.NewVBox(
		bar,
		name,
		cancel,
		delete,
	))
}

func settingsWin(fynewindow fyne.Window) {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin(fynewindow)
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin(fynewindow)
		}),
	)

	cols := widget.NewEntry()
	cols.SetPlaceHolder("Enter the number of columns you want there to be here")

	themes := widget.NewEntry()
	themes.SetPlaceHolder("Enter the theme you want to use here (1 = dark, 2 = light)")

	cancel := widget.NewButton("Cancel", func() {
		mainWin(fynewindow)
	})

	finish := widget.NewButton("Finish", func() {
		settings := openSettings()

		col, err := strconv.Atoi(cols.Text)
		if err != nil {
			col = settings.Columns
		}

		themeInt, err := strconv.Atoi(themes.Text)
		if err != nil {
			themeInt = settings.Theme
		}
		switch themeInt {
		case 1:
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		case 2:
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		default:
			log.Println("Invalid theme")
			return
		}

		writeSettings(Settings{Columns: col, Theme: themeInt})

		mainWin(fynewindow)
	})

	fynewindow.SetContent(container.NewVBox(
		bar,
		cols,
		themes,
		cancel,
		finish,
	))
}

func mainWin(fynewindow fyne.Window) {
	buttons := jsonToBtns(fynewindow)
	fynewindow.SetContent(container.NewVBox(
		buttons...,
	))
}

func main() {
	a := app.NewWithID("com.minetest-j45.go_soundboard")
	a.SetIcon(gosoundboardlogoPng)

	w := a.NewWindow("Go Soundboard - 0.0")

	mainWin(w)

	w.Resize(fyne.NewSize(700, 400))
	w.ShowAndRun()
}
