package scanner

import (
	"context"
	"fmt"
	"os"
	"railsearch/pkg/config"
	"railsearch/pkg/handler"
	"time"

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
	rScan.Handler.SetConfig(rScan.Config)

	skip := rScan.Handler.GetSkips()

	if skip.SkipNode {
		scanner.SkipNodes = true
	}

	if skip.SkipRelation {
		scanner.SkipRelations = true
	}

	if skip.SkipWay {
		scanner.SkipWays = true
	}

	go rScan.Handler.HandleRelation(relationchan)
	go rScan.Handler.HandleNode(nodechan)
	go rScan.Handler.HandleWay(waychan)

	fmt.Println("BEGINNING SCANNING BEEP")
	beginTime := time.Now()

	i := 0
	for scanner.Scan() {
		i = i + 1

		if i%10000 == 0 {
			fmt.Println("searching ", i, "objects")
		}
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

	fmt.Println("DONE SCANNING, THANK ME BITCH")
	fmt.Println("Scan begins at: ", beginTime)
	fmt.Println("Scan ends   at: ", time.Now())
	fmt.Println("Time   elapsed: ", time.Since(beginTime))

	close(nodechan)
	close(waychan)
	close(relationchan)
}
