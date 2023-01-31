package main

import (
	"railsearch/pkg/config"
	"railsearch/pkg/handler"
	"railsearch/pkg/scanner"
)

func main() {
	config := config.GetConfig("config.json")

	scanner := scanner.NewScanner(config, handler.NewNodeHandler())

	scanner.Scan()

}
