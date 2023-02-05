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

	scan_search := scanner.NewScanner(config, handler.NewSearchHandler())
	scan_search.Scan()

	// scan_way := scanner.NewScanner(config, handler.NewWayHandler())
	// scan_way.Scan()

	// scan_building := scanner.NewScanner(config, handler.NewBuildingHandler())
	// scan_building.Scan()

}
