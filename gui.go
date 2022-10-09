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

var fynewindow fyne.Window

func jsonToBtns() []fyne.CanvasObject {
	buttons := openJson()

	var btns []fyne.CanvasObject

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin()
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin()
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

func addWin() {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin()
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin()
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter new sound name here")
	file := widget.NewEntry()
	file.SetPlaceHolder("Enter new sound file here")

	cancel := widget.NewButton("Cancel", func() {
		mainWin()
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
			mainWin()
			confNewSound(name.Text, file.Text)
			mainWin()
		case ".mp3":
			mainWin()
			confNewSound(name.Text, file.Text)
			mainWin()
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

func rmWin() {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin()
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin()
		}),
	)

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter the name of the sound you want to delete here")

	cancel := widget.NewButton("Cancel", func() {
		mainWin()
	})

	delete := widget.NewButton("Delete", func() {
		confDeleteSound(name.Text)
	})

	fynewindow.SetContent(container.NewVBox(
		bar,
		name,
		cancel,
		delete,
	))
}

func settingsWin() {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			mainWin()
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			addWin()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			rmWin()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsWin()
		}),
	)

	cols := widget.NewEntry()
	cols.SetPlaceHolder("Enter the number of columns you want there to be here")

	themes := widget.NewEntry()
	themes.SetPlaceHolder("Enter the theme you want to use here (1 = dark, 2 = light)")

	cancel := widget.NewButton("Cancel", func() {
		mainWin()
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

		mainWin()
	})

	fynewindow.SetContent(container.NewVBox(
		bar,
		cols,
		themes,
		cancel,
		finish,
	))
}

func mainWin() {
	buttons := jsonToBtns()
	fynewindow.SetContent(container.NewVBox(
		buttons...,
	))
}

func main() {
	a := app.NewWithID("com.minetest-j45.go_soundboard")
	a.SetIcon(gosoundboardlogoPng)

	fynewindow = a.NewWindow("Go Soundboard - 0.0")

	mainWin()

	fynewindow.Resize(fyne.NewSize(700, 400))
	fynewindow.ShowAndRun()
}
