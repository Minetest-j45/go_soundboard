/*
Soundboard is a simple soundboard program
made for learning purposes

Usage:
	soundboard
*/
package main

import (
	"github.com/gordonklaus/portaudio"
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	w := helloWindow()
	w.ShowAndRun()
}
