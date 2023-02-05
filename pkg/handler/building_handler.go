package handler

import (
	"fmt"
	"railsearch/pkg/config"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/project"
	"github.com/paulmach/osm"
)

type BuildingHandler struct {
	Conf config.RailsearchConfig
}

func NewBuildingHandler() Handler {
	return &BuildingHandler{}
}

type Coordinate struct {
	Longitude  float64
	Latitude   float64
	CartesianX float64
	CartesianY float64
}

func (h *BuildingHandler) HandleNode(channel chan *osm.Node) {
	i := 0
	for _ = range channel {
		i = i + 1
		// if h.validTags(b.Tags) {

		// mercator := project.Point(orb.Point{b.Lon, b.Lat}, project.WGS84.ToMercator)
		// coorx := mercator.X()
		// coory := mercator.Y()

		// x1 := coorx - float64(h.Conf.Range)
		// x2 := coorx + float64(h.Conf.Range)
		// y1 := coory - float64(h.Conf.Range)
		// y2 := coory + float64(h.Conf.Range)
		// var node database.Node

		// // building := database.BuildingNode{
		// // 	OsmId:     int64(b.ID),
		// // 	Latitude:  b.Lat,
		// // 	Longitude: b.Lon,
		// // 	OsmType:   "Node",
		// // }
		// err := database.GetConn().Unscoped().Where("cartesian_x BETWEEN ? AND ?", x1, x2).Where("cartesian_y BETWEEN ? AND ?", y1, y2).First(&node).Error
		// if err == nil {
		// 	i = i + 1
		// 	// fmt.Println(building)
		// 	// database.GetConn().Create(&building)
		// } else {
		// 	var way database.Way
		// 	err := database.GetConn().Unscoped().Where("cartesian_x BETWEEN ? AND ?", x1, x2).Where("cartesian_y BETWEEN ? AND ?", y1, y2).First(&way).Error
		// 	if err == nil {
		// 		i = i + 1
		// 		// fmt.Println(building)
		// 		// database.GetConn().Create(&building)
		// 	}
		// }
		// }
	}

	fmt.Println("Theres", i, "Nodes that have amenity and building tag")
}

func (h *BuildingHandler) HandleWay(channel chan *osm.Way) {
	i := 0
	for _ = range channel {
		// if h.validTags(b.Tags) {

		// 	coordinate := findCentroid(*b)

		// 	x1 := coordinate.CartesianX - float64(h.Conf.Range)
		// 	x2 := coordinate.CartesianX + float64(h.Conf.Range)
		// 	y1 := coordinate.CartesianY - float64(h.Conf.Range)
		// 	y2 := coordinate.CartesianY + float64(h.Conf.Range)
		// 	var node database.Node

		// 	building := database.BuildingNode{
		// 		OsmId:     int64(b.ID),
		// 		Latitude:  coordinate.Latitude,
		// 		Longitude: coordinate.Longitude,
		// 		OsmType:   "Way",
		// 	}
		// 	fmt.Println(building)

		// 	err := database.GetConn().Unscoped().Where("cartesian_x BETWEEN ? AND ?", x1, x2).Where("cartesian_y BETWEEN ? AND ?", y1, y2).First(&node).Error
		// 	if err == nil {
		// 		i = i + 1
		// 		fmt.Println(building)
		// 		// database.GetConn().Create(&building)
		// 	} else {
		// 		var way database.Way
		// 		err := database.GetConn().Unscoped().Where("cartesian_x BETWEEN ? AND ?", x1, x2).Where("cartesian_y BETWEEN ? AND ?", y1, y2).First(&way).Error
		// 		if err == nil {
		// 			i = i + 1
		// 			fmt.Println(building)
		// 			// database.GetConn().Create(&building)
		// 		}
		// 	}
		// }
	}

	fmt.Println("Theres", i, "ways that have amenity and building tag")
}

func (h *BuildingHandler) HandleRelation(channel chan *osm.Relation) {
}

func (h *BuildingHandler) SetConfig(conf config.RailsearchConfig) {
	h.Conf = conf
}

func (h *BuildingHandler) GetSkips() SkipObject {
	return SkipObject{
		SkipRelation: true,
		SkipWay:      true,
		SkipNode:     false,
	}
}

func (h *BuildingHandler) validTags(tags osm.Tags) bool {
	for _, tc := range h.Conf.AroundTag {
		found := tags.Find(tc.TagName)
		if found != "" {
			if len(tc.Blacklist) == 0 && len(tc.Whitelist) == 0 {
				return true
			}
			if len(tc.Blacklist) > 0 {
				for _, val := range tc.Blacklist {
					if val == found {
						return false
					}
				}
				return true
			}

			if len(tc.Whitelist) > 0 {
				for _, val := range tc.Whitelist {
					if val == found {
						return true
					}
				}
				return false
			}
		}
	}
	return false
}

func findCentroid(way osm.Way) Coordinate {

	totalLat := 0.0
	totalLon := 0.0

	for _, node := range way.Nodes {
		totalLat = totalLat + node.Lat
		totalLon = totalLon + node.Lon
	}

	centrolat := totalLat / float64(len(way.Nodes))
	centrolon := totalLon / float64(len(way.Nodes))

	mercator := project.Point(orb.Point{centrolon, centrolat}, project.WGS84.ToMercator)
	return Coordinate{
		Latitude:   centrolat,
		Longitude:  centrolon,
		CartesianX: mercator.X(),
		CartesianY: mercator.Y(),
	}
}
