package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetConfig() Config {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error occured while reading config")
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error:", err)
	}

	return config
}