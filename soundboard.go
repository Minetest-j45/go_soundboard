/*
Soundboard is a simple soundboard program
made for learning purposes

Usage:
	soundboard
*/
package main

import (
	"log"
)

func main() {
	w := mainWindow()
	w.ShowAndRun()
	
	for {
		log.Println(showNewSoundWindow)
		if showNewSoundWindow == true {
			w := newSoundWindow()
			w.ShowAndRun()
			showNewSoundWindow = false
		}
	}
}
