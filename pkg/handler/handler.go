package handler

import (
	"railsearch/pkg/config"

	"github.com/paulmach/osm"
)

type Handler interface {
	SetConfig(config.RailsearchConfig)
	GetSkips() SkipObject
	HandleNode(chan *osm.Node)
	HandleWay(chan *osm.Way)
	HandleRelation(chan *osm.Relation)
}

type SkipObject struct {
	SkipRelation bool
	SkipWay      bool
	SkipNode     bool
}
