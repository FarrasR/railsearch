package handler

import (
	"fmt"
	"railsearch/pkg/config"
	"railsearch/pkg/database"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/project"
	"github.com/paulmach/osm"
)

type TargetHandler struct {
	Conf config.RailsearchConfig
}

func NewTargetHandler() Handler {
	return &TargetHandler{}
}

func (h *TargetHandler) HandleNode(channel chan *osm.Node) {
	i := 0
	for b := range channel {
		if h.validTags(b.Tags) {
			mercator := project.Point(orb.Point{b.Lon, b.Lat}, project.WGS84.ToMercator)
			coorx := mercator.X()
			coory := mercator.Y()

			x1 := coorx - float64(h.Conf.Range)
			x2 := coorx + float64(h.Conf.Range)
			y1 := coory - float64(h.Conf.Range)
			y2 := coory + float64(h.Conf.Range)

			var indexNode database.IndexNode

			err := database.GetConn().Unscoped().
				Where("cartesian_x BETWEEN ? AND ?", x1, x2).
				Where("cartesian_y BETWEEN ? AND ?", y1, y2).
				First(&indexNode).Error

			if err == nil {
				//this is a node that near the index node
				i = i + 1

				targetNode := database.TargetNode{
					OsmId:     int64(b.ID),
					Latitude:  b.Lat,
					Longitude: b.Lon,
					Tags:      fmt.Sprintf("%v", b.Tags),
					Type:      "node",
				}

				// fmt.Println(targetNode)
				database.GetConn().Create(&targetNode)
			}
		}
	}
	fmt.Println("theres", i, "node objects near target")
}

func (h *TargetHandler) HandleWay(channel chan *osm.Way) {
	i := 0
	for b := range channel {
		if h.validTags(b.Tags) {
			var node = database.Node{NodeId: int64(b.Nodes[0].ID)}
			database.GetConn().First(&node)

			mercator := project.Point(orb.Point{node.Longitude, node.Latitude}, project.WGS84.ToMercator)
			coorx := mercator.X()
			coory := mercator.Y()

			x1 := coorx - float64(h.Conf.Range)
			x2 := coorx + float64(h.Conf.Range)
			y1 := coory - float64(h.Conf.Range)
			y2 := coory + float64(h.Conf.Range)

			var indexNode database.IndexNode

			err := database.GetConn().Unscoped().
				Where("cartesian_x BETWEEN ? AND ?", x1, x2).
				Where("cartesian_y BETWEEN ? AND ?", y1, y2).
				First(&indexNode).Error

			if err == nil {
				//this is a node that near the index node
				i = i + 1

				targetNode := database.TargetNode{
					OsmId:     int64(b.ID),
					Latitude:  node.Latitude,
					Longitude: node.Longitude,
					Tags:      fmt.Sprintf("%v", b.Tags),
					Type:      "way",
				}

				// fmt.Println(targetNode)
				database.GetConn().Create(&targetNode)
			}

		}
	}
	fmt.Println("theres", i, "way objects near target")
}

func (h *TargetHandler) HandleRelation(channel chan *osm.Relation) {
}

func (h *TargetHandler) SetConfig(conf config.RailsearchConfig) {
	h.Conf = conf
}

func (h *TargetHandler) GetSkips() SkipObject {
	return SkipObject{
		SkipRelation: true,
		SkipWay:      false,
		SkipNode:     false,
	}
}

type Coordinate struct {
	Longitude  float64
	Latitude   float64
	CartesianX float64
	CartesianY float64
}

//centroid function that might not be used
//very costly on database
func FindCentroid(way osm.Way) Coordinate {

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

func (h *TargetHandler) validTags(tags osm.Tags) bool {
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
