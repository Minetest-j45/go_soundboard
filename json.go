package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
)

type Buttons struct {
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Name string `json:"name"`
	File string `json:"file"`
}

func openJson() Buttons {
	jsonFile, err := os.Open("./soundboard.json")
	if err != nil {
		os.Create("./soundboard.json")
		openJson()
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var buttons Buttons

	json.Unmarshal(byteValue, &buttons)

	return buttons
}

func confNewSound(name string, file string) {
	buttons := openJson()

	newButton := Button{}

	newButton.Name = name
	newButton.File = file

	buttons.Buttons = append(buttons.Buttons, newButton)

	newButtonBytes, err := json.MarshalIndent(buttons, "", " ")
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile("./soundboard.json", newButtonBytes, 0666)
}

func confDeleteSound(name string, fynewindow fyne.Window) {
	buttons := openJson()

	for _, v := range buttons.Buttons {
		if v.Name == name {
			length := len(buttons.Buttons)
			for index, field := range buttons.Buttons {
				if field.Name == name {
					if index == length-1 {
						buttons.Buttons = buttons.Buttons[0:index]
					} else {
						buttons.Buttons = append(buttons.Buttons[0:index], buttons.Buttons[index+1:]...)
					}
				}
			}

			out, _ := json.MarshalIndent(buttons, "", "  ")
			_ = ioutil.WriteFile("./soundboard.json", out, 0666)
			mainWindowSetContext(fynewindow)
		}
	}
}

func confExists(name string) bool {
	buttons := openJson()

	var exists bool = false
	for _, button := range buttons.Buttons {
		if button.Name == name {
			exists = true
		}
	}

	return exists
}
