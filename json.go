package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

type Buttons struct {
    Buttons []Button `json:"buttons"`
}

type Button struct {
    Name   string `json:"name"`
    File   string `json:"file"`
    Number   int `json:"number"`
}

func openJson(file string) {
    jsonFile, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened " + file)
    defer jsonFile.Close()
  
    byteValue, _ := ioutil.ReadAll(jsonFile)
  
    var buttons Buttons
  
    json.Unmarshal(byteValue, &buttons)
  
    for i := 0; i < len(buttons.Buttons); i++ {
        fmt.Println("Button name: " + buttons.Buttons[i].Name)
        fmt.Println("Button file: " + buttons.Buttons[i].File)
        fmt.Println("Button number: " + strconv.Itoa(buttons.Buttons[i].Number))
    }
}
