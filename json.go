package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"strconv"
)

type Buttons struct {
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Name   string `json:"name"`
	File   string `json:"file"`
	Number int    `json:"number"`
}

func openJson() Buttons {
	jsonFile, err := os.Open("./soundboard.json")
	if err != nil {
		os.Create("./soundboard.json")
		openJson()
	}
	fmt.Println("Successfully Opened soundboard.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var buttons Buttons

	json.Unmarshal(byteValue, &buttons)

	/*for i := 0; i < len(buttons.Buttons); i++ {
		fmt.Println("Button name: " + buttons.Buttons[i].Name)
		fmt.Println("Button file: " + buttons.Buttons[i].File)
		fmt.Println("Button number: " + strconv.Itoa(buttons.Buttons[i].Number))
	}*/
	return buttons
}

func confNewSound(name string, file string) {
	jsonFile, err := os.Open("./soundboard.json")
	if err != nil {
		os.Create("./soundboard.json")
		confNewSound(name, file)
	}
	fmt.Println("Successfully Opened `./soundboard.json`")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var buttons Buttons

	json.Unmarshal(byteValue, &buttons)

	newButton := Button{}

	newButton.Name = name
	newButton.File = file
	newButton.Number = len(buttons.Buttons)

	buttons.Buttons = append(buttons.Buttons, newButton)

	newButtonBytes, err1 := json.MarshalIndent(buttons, "", " ")
	if err1 != nil {
		log.Println(err1)
	}

	ioutil.WriteFile("./soundboard.json", newButtonBytes, 0666)
}

/*func confDeleteSound(name string) {
	//read json file, find the name of the sound, delete it
	//if it was in json file, main window context
	//else log invalid sound name
}*/
