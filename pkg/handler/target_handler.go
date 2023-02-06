package handler

import (
	"railsearch/pkg/config"

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
}

func (h *TargetHandler) HandleWay(channel chan *osm.Way) {
}

func (h *TargetHandler) HandleRelation(channel chan *osm.Relation) {
}

func (h *TargetHandler) SetConfig(conf config.RailsearchConfig) {
	h.Conf = conf
}

func (h *TargetHandler) GetSkips() SkipObject {
	return SkipObject{
		SkipRelation: true,
		SkipWay:      true,
		SkipNode:     false,
	}
}

type Coordinate struct {
	Longitude  float64
	Latitude   float64
	CartesianX float64
	CartesianY float64
}

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
