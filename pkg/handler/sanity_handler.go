package handler

import (
	"fmt"
	"railsearch/pkg/config"

	"github.com/paulmach/osm"
)

type SanityHandler struct {
}

func NewSanityHandler() Handler {
	return &SanityHandler{}
}

func (h *SanityHandler) HandleNode(channel chan *osm.Node) {
	i := 0
	for range channel {
		i = i + 1
	}
	fmt.Println("theres", i, "node objects")
}

func (h *SanityHandler) HandleWay(channel chan *osm.Way) {
	i := 0
	for range channel {
		i = i + 1
	}
	fmt.Println("theres", i, "way objects")
}

func (h *SanityHandler) HandleRelation(channel chan *osm.Relation) {
}

func (h *SanityHandler) SetConfig(conf config.RailsearchConfig) {
}

func (h *SanityHandler) GetSkips() SkipObject {
	return SkipObject{
		SkipRelation: true,
		SkipWay:      false,
		SkipNode:     false,
	}
}
