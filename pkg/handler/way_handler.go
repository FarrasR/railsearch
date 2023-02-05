package handler

import (
	"railsearch/pkg/config"

	"github.com/paulmach/osm"
)

type WayHandler struct {
	Conf config.RailsearchConfig
}

func NewWayHandler() Handler {
	return &WayHandler{}
}

func (h *WayHandler) HandleNode(channel chan *osm.Node) {
	for _ = range channel {
		// var wayNode database.WayMember
		// if err := database.GetConn().Unscoped().First(&wayNode, "node_id = ?", int64(b.ID)).Error; err == nil {
		// 	way := database.Way{
		// 		WayId:     int64(b.ID),
		// 		Latitude:  b.Lat,
		// 		Longitude: b.Lon,
		// 	}
		// 	database.GetConn().Create(&way)
		// }
	}
}

func (h *WayHandler) HandleWay(channel chan *osm.Way) {
}

func (h *WayHandler) HandleRelation(channel chan *osm.Relation) {
}

func (h *WayHandler) SetConfig(conf config.RailsearchConfig) {
	h.Conf = conf
}

func (h *WayHandler) GetSkips() SkipObject {
	return SkipObject{
		SkipRelation: true,
		SkipWay:      true,
		SkipNode:     false,
	}
}
