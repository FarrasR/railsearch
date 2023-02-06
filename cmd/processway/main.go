package main

import (
	"fmt"
	"railsearch/pkg/config"
	"railsearch/pkg/database"
	"time"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/project"
)

func main() {
	config := config.GetConfig("config.json")
	database.InitDB(config.DatabaseConfig)

	rows, err := database.GetConn().Model(&database.WayNode{}).Rows()
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	beginTime := time.Now()
	fmt.Println("Starting converting at", beginTime)

	i := 0
	for rows.Next() {
		i = i + 1

		if i%100 == 0 {
			fmt.Println("converting ", i, " ways")
		}
		var way database.WayNode
		database.GetConn().ScanRows(rows, &way)

		var node = database.Node{
			NodeId: way.NodeId,
		}
		database.GetConn().First(&node)

		mercator := project.Point(orb.Point{node.Longitude, node.Latitude}, project.WGS84.ToMercator)
		indexNode := database.IndexNode{
			NodeId:     node.NodeId,
			CartesianX: mercator.X(),
			CartesianY: mercator.Y(),
		}
		database.GetConn().Create(&indexNode)
	}

	fmt.Println("Conversion begins at:", beginTime)
	fmt.Println("Conversion ends   at:", time.Now())
	fmt.Println("Time         elapsed:", time.Since(beginTime))

}
