/*
Soundboard is a simple soundboard program
made for learning purposes

Usage:
	soundboard
*/
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/storage"
	"log"
)

func main() {
	w := mainWindow()
	w.ShowAndRun()
}
