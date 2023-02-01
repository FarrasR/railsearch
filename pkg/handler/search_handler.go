package handler

import (
	"railsearch/pkg/config"
	"railsearch/pkg/database"

	"github.com/paulmach/osm"
)

type SearchHandler struct {
	Conf config.RailsearchConfig
}

func NewSearchHandler() Handler {
	return &SearchHandler{}
}

func (h *SearchHandler) HandleNode(channel chan *osm.Node) {
	for b := range channel {
		if h.validTags(b.Tags) {
			node := database.Node{
				NodeId:    int64(b.ID),
				Latitude:  b.Lat,
				Longitude: b.Lon,
			}
			database.GetConn().Create(&node)
		}
	}
}

func (h *SearchHandler) HandleWay(channel chan *osm.Way) {
	for b := range channel {
		if h.validTags(b.Tags) {
			for i, val := range b.Nodes {
				way := database.WayMember{
					WayId:  int64(b.ID),
					NodeId: int64(val.ID),
					Order:  i,
				}
				database.GetConn().Create(&way)
			}
		}
	}
}

func (h *SearchHandler) HandleRelation(channel chan *osm.Relation) {
}

func (h *SearchHandler) SetConfig(conf config.RailsearchConfig) {
	h.Conf = conf
}

func (h *SearchHandler) GetSkips() SkipObject {
	return SkipObject{
		SkipRelation: true,
		SkipWay:      false,
		SkipNode:     false,
	}
}

func (h *SearchHandler) validTags(tags osm.Tags) bool {
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
