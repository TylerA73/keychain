package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	r "math/rand"
	"time"
)

var config Settings

type Settings struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}

func init() {
	ImportSettings()
}

/**
 * Generate Key
 */
func GenerateKey(length int) string {
	r.Seed(time.Now().UnixNano())
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()")
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

func CreateFile() Settings {
	fmt.Println("File being created...")
	data := Settings{
		Key:  GenerateKey(20),
		Hash: CreateHash(GenerateKey(32)),
	}
	file, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile("./settings/config.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func ImportSettings() {
	file, err := ioutil.ReadFile("./settings/config.json")
	if err != nil {
		config = CreateFile()
	} else {
		_ = json.Unmarshal([]byte(file), &config)
	}
}
