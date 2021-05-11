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
	w.SetContent(taskapp.makeUI())
	w.ShowAndRun()
}
