/*package main

import (
	"log"
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newSoundWindow(fyneapp fyne.App) fyne.Window {
    s := fyneapp.NewWindow("New Sound - Go Soundboard")

    hello := widget.NewLabel("Hello, World!")
    input := widget.NewEntry()
    input.SetPlaceHolder("Enter new sound name here")
    content := container.NewVBox(input, widget.NewButton("Save", func() {
        log.Println("Content was:", input.Text)
    }))
    s.SetContent(container.NewVBox(
        hello,
        content,
    ))

    return s
}

func mainWindow() fyne.Window {
    a := app.New()
    w := a.NewWindow("Go Soundboard")

    hello := widget.NewLabel("Hello, World!")
    w.SetContent(container.NewVBox(
        hello,
        widget.NewButton("Hi", func() {
            hello.SetText("Recording")
            recordAudio()
            hello.SetText("Playing")
            // playAudio(audio)
        }),
        widget.NewButton("+", func() {
            hello.SetText("Making a new sound")
            newSoundWindow(a)
            hello.SetText("Hello, World!")
        }),
    ))

    return w
}*/
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type taskApp struct {
	app             fyne.App
	win             fyne.Window
	settings_uri    fyne.URI /* file with settings */

	helpWindow     fyne.Window
	settingsWindow fyne.Window
}

var about_text = `some text here`

var help_text = `some text here`

/*****************************************************************************/
/*                         Create Menu                                       */
/*****************************************************************************/
func create_menu(a *taskApp) *fyne.MainMenu {

	darkItem := fyne.NewMenuItem("Dark Theme", func() {
		a.app.Settings().SetTheme(theme.DarkTheme())
	})

	lightItem := fyne.NewMenuItem("Light Theme", func() {
		a.app.Settings().SetTheme(theme.LightTheme())
	})


	aboutItem := fyne.NewMenuItem("About", func() {
		dialog.ShowInformation("About", about_text, a.win)
	})

	helpItem := fyne.NewMenuItem("Help", func() {
		a.helpWindow = a.app.NewWindow("Help")
		a.helpWindow.SetContent(a.make_help())
		a.helpWindow.Resize(fyne.NewSize(480, 720))
		a.helpWindow.Show()
	})

	return fyne.NewMainMenu(
		fyne.NewMenu("Settings", lightItem, darkItem, fyne.NewMenuItemSeparator()),
		fyne.NewMenu("Help", aboutItem, helpItem),
	)
}

/*****************************************************************************/
/*                         Create New Tab                                    */
/*****************************************************************************/
func create_new_tab(a *taskApp) *container.Scroll {
	bar_details := widget.NewToolbar(
		// save task
		widget.NewToolbarAction(theme.ConfirmIcon(), func() {
			a.f_stored = false
			a.SaveDetails()
		}),
	)

	a.w_ID = widget.NewLabel("")
	a.w_modtime = widget.NewLabel("")
	a.w_name = widget.NewEntry()
	a.w_description = widget.NewEntry()
	a.w_sound_file = widget.NewEntry() /
	details := widget.NewForm(
		widget.NewFormItem("Name", a.w_name),
		widget.NewFormItem("Description", a.w_description),
		widget.NewFormItem("Sound file", a.w_sound_file),
	)

	a.w_new_done = widget.NewCheck("Done", func(v bool) {})

	box_details := container.NewScroll(container.NewVBox(
		widget.NewLabel(""),
		container.NewHBox(
			//widget.NewLabel ("save: "),
			bar_details,
		),
		container.New(layout.NewCenterLayout(), widget.NewLabel("New")),
		container.NewHBox(widget.NewLabel("ID:       "), a.w_ID),
		container.NewHBox(widget.NewLabel("Mod-Time: "), a.w_modtime),
		details,
		a.w_details_done,
	))
	return box_details
}

/* =============================================================================================== */
func (a *taskApp) makeUI() fyne.CanvasObject {

	/* Create Menu */
	a.win.SetMainMenu(create_menu(a))
	
	/* new tab */
	box_new := create_new_tab(a)
	
	a.w_tab_new = container.NewTabItem("New", box_new)

	a.tabbar = container.NewAppTabs()
	a.tabbar.OnChanged = func(item *container.TabItem) {
		if item == a.w_tab_new {
			a.DisplayCurrentTask()
		}
	}

	return a.tabbar
}



/* =============================================================================================== */
func (a *taskApp) make_help() fyne.CanvasObject {

	e := widget.NewMultiLineEntry()
	e.Text = help_text
	e.Disable()
	e.Wrapping = fyne.TextWrapWord
	return e
}
