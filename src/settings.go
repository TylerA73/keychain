package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var config Settings

type Settings struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}

func init() {
	ImportSettings()
}

func CreateFile() Settings {
	data := Settings{
		Key: CreateHash("testingthisout"),
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("../settings/config.json", file, 0644)
	return data
}

func ImportSettings() {
	file, err := ioutil.ReadFile("../settings/config.json")
	if err != nil {
		config = CreateFile()
	} else {
		_ = json.Unmarshal([]byte(file), &config)
	}

	fmt.Println("Key: ", config.Key)
}
