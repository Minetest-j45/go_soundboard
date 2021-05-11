/*
Soundboard is a simple soundboard program
made for learning purposes

Usage:
	soundboard
*/
package main

func main() {
	w := mainWindow()
	w.ShowAndRun()
	
	for {
		if showNewSoundWindow == true {
			fmt.Println("showNewSoundWindow = true")
			w := newSoundWindow()
			w.ShowAndRun()
			showNewSoundWindow = false
		}
	}
}
