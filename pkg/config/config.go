package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type RailsearchConfig struct {
	FileName  string      `json:"file_name"`
	SearchTag []TagConfig `json:"search_tag"`
	AroundTag []TagConfig `json:"around_tag"`
	Range     int         `json:"range"`
}

type TagConfig struct {
	TagName   string   `json:"name"`
	Blacklist []string `json:"blacklist"`
	Whitelist []string `json:"whitelist"`
}

func GetConfig(file string) RailsearchConfig {
	config_file, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer config_file.Close()

	byteValue, _ := ioutil.ReadAll(config_file)

	var config RailsearchConfig

	json.Unmarshal(byteValue, &config)

	return config
}
