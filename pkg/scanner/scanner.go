package scanner

import (
	"context"
	"os"
	"railsearch/pkg/config"
	"railsearch/pkg/handler"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
)

type RailsearchScanner struct {
	Config  config.RailsearchConfig
	Handler handler.Handler
	Scanner *osmpbf.Scanner
}

func NewScanner(Config config.RailsearchConfig, Handler handler.Handler) *RailsearchScanner {
	return &RailsearchScanner{
		Config:  Config,
		Handler: Handler,
	}
}

func (rScan *RailsearchScanner) Scan() {
	f, err := os.Open(rScan.Config.FileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()
	nodechan := make(chan *osm.Node)
	waychan := make(chan *osm.Way)
	relationchan := make(chan *osm.Relation)

	go rScan.Handler.HandleRelation(relationchan)
	go rScan.Handler.HandleNode(nodechan)
	go rScan.Handler.HandleWay(waychan)

	for scanner.Scan() {
		object := scanner.Object()
		switch object.(type) {
		case *osm.Node:
			nodechan <- object.(*osm.Node)
		case *osm.Way:
			waychan <- object.(*osm.Way)
		case *osm.Relation:
			relationchan <- object.(*osm.Relation)
		}
	}
}
