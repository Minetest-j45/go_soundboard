package main

import (
  "encoding/json"
)

type Buttons struct {
    Buttons []Button `json:"buttons"`
}

type Button struct {
    Name   string `json:"name"`
    Type   string `json:"file"`
    Number   string `json:"number"`
}

func openJson(file string) {
  jsonFile, err := os.Open(file)
  if err != nil {
    fmt.Println(err)
  }
  defer jsonFile.Close()
}
