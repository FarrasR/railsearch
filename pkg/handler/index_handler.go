package handler

import (
	"railsearch/pkg/config"
	"railsearch/pkg/database"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/project"
	"github.com/paulmach/osm"
)

type IndexHandler struct {
	Conf config.RailsearchConfig
}

func NewSearchHandler() Handler {
	return &IndexHandler{}
}

func (h *IndexHandler) HandleNode(channel chan *osm.Node) {
	for b := range channel {
		node := database.Node{
			NodeId:    int64(b.ID),
			Latitude:  b.Lat,
			Longitude: b.Lon,
		}
		database.GetConn().Create(&node)
		if h.validTags(b.Tags) {
			mercator := project.Point(orb.Point{b.Lon, b.Lat}, project.WGS84.ToMercator)
			indexNode := database.IndexNode{
				NodeId:     int64(b.ID),
				CartesianX: mercator.X(),
				CartesianY: mercator.Y(),
			}
			database.GetConn().Create(&indexNode)
		}
	}
}

func (h *IndexHandler) HandleWay(channel chan *osm.Way) {
	for b := range channel {
		if h.validTags(b.Tags) {
			for i, val := range b.Nodes {
				way := database.WayNode{
					WayId:  int64(b.ID),
					NodeId: int64(val.ID),
					Order:  i,
				}
				database.GetConn().Create(&way)
			}
		}
	}
}

func (h *IndexHandler) HandleRelation(channel chan *osm.Relation) {
}

func (h *IndexHandler) SetConfig(conf config.RailsearchConfig) {
	h.Conf = conf
}

func (h *IndexHandler) GetSkips() SkipObject {
	return SkipObject{
		SkipRelation: true,
		SkipWay:      true,
		SkipNode:     false,
	}
}

func (h *IndexHandler) validTags(tags osm.Tags) bool {
	for _, tc := range h.Conf.SearchTag {
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
