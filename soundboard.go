/*
Soundboard is a simple soundboard program
made for learning purposes

Usage:
	soundboard
*/
package main

func main() {
	a := app.New()
	w := a.NewWindow("Go Soundboard")
	
	uri1, err1 := storage.Child(app.Storage().RootURI(), "settings.json")
	if err1 != nil {
		log.Println("tasks_main: error-1: ", err1)
	} else {
		log.Println("tasks_main: URI-2-String: ", uri1.String())
	}

	
	taskapp := &taskApp{
		app:             a,
		win:             w,
		settings_uri:    uri1,
	}
	
	taskapp.readSettings(uri1)
	
	w.SetContent(taskapp.makeUI())
	w.Resize(fyne.NewSize(400, 700))
	w.ShowAndRun()
}
