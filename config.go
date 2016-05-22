package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
)

// Config is application config
type Config struct {
	Counters map[string]int64 `json:"counters"`
}

// newConfig is config.json parse
func newConfig() *Config {
	return &Config{
		Counters: map[string]int64{
			"default": 0,
		},
	}
}

// marshalJSON is config.json parse to json
func (c *Config) marshalJSON() {
	output, err := json.MarshalIndent(&c, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	// configファイルのディレクトリ取得
	directory := getConfigDirectory()
	err = ioutil.WriteFile(directory, output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

// unmarshalJSON is config.json parse to struct
func unmarshalJSON() *Config {
	// configファイルのディレクトリ取得
	directory := getConfigDirectory()
	// ファイルの読み込み
	rf, err := ioutil.ReadFile(directory) //[]byte型での読み込み
	if err != nil {
		log.Fatal(err)
	}
	data := new(Config)
	if err := json.Unmarshal(rf, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}
	return data
}

func isConfigDirectoryExist() bool {
	// configファイルのディレクトリ取得
	directory := getConfigDirectory()
	initFlag := isExist(directory)
	return initFlag
}
func getConfigDirectory() string {
	homedir, _ := homedir.Dir()
	configDirectory := homedir + "/" + CONFIG
	return configDirectory
}
