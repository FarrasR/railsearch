package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"railsearch/pkg/config"
)

func main() {
	fmt.Println("opening json file...")
	config_file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer config_file.Close()

	byteValue, _ := ioutil.ReadAll(config_file)

	var config config.RailsearchConfig

	json.Unmarshal(byteValue, &config)

	fmt.Println(config)
}
