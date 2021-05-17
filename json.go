package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Buttons struct {
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Name   string `json:"name"`
	File   string `json:"file"`
	Number int    `json:"number"`
}

func openJson(file string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + file)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var buttons Buttons

	json.Unmarshal(byteValue, &buttons)

	for i := 0; i < len(buttons.Buttons); i++ {
		fmt.Println("Button name: " + buttons.Buttons[i].Name)
		fmt.Println("Button file: " + buttons.Buttons[i].File)
		fmt.Println("Button number: " + strconv.Itoa(buttons.Buttons[i].Number))
	}
}

/*func confNewSound(name string, file string) {
	jsonFile, err := os.Open("./soundboard.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened `./soundboard.json`")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var buttons Buttons

	json.Unmarshal(byteValue, &buttons)
	
	buttons.Buttons[len(buttons.Buttons)+1].Name := name
	buttons.Buttons[len(buttons.Buttons)+1].File := file
	buttons.Buttons[len(buttons.Buttons)+1].Number := strconv.Itoa(len(buttons.Buttons)+1)
}*/
