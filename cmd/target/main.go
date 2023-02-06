package main

import (
	"railsearch/pkg/config"
	"railsearch/pkg/database"
	"railsearch/pkg/handler"
	"railsearch/pkg/scanner"
)

func main() {
	config := config.GetConfig("config.json")
	database.InitDB(config.DatabaseConfig)

	scan_search := scanner.NewScanner(config, handler.NewTargetHandler())
	scan_search.Scan()
}
