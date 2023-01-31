package handler

import (
	"fmt"

	"github.com/paulmach/osm"
)

type Handler interface {
	HandleNode(chan *osm.Node)
	HandleWay(chan *osm.Way)
	HandleRelation(chan *osm.Relation)
}

type NodeHandler struct {
}

func NewNodeHandler() Handler {
	return &NodeHandler{}
}

func (*NodeHandler) HandleNode(channel chan *osm.Node) {
	for b := range channel {
		if b == nil {
			break
		}
		fmt.Println(b.Tags)
	}
	fmt.Println("wadidau")
}

func (*NodeHandler) HandleWay(channel chan *osm.Way) {
	for b := range channel {
		if b == nil {
			break
		}
		fmt.Println(b.Tags)
	}
}

func (*NodeHandler) HandleRelation(channel chan *osm.Relation) {
	for b := range channel {
		if b == nil {
			break
		}
		fmt.Println(b.Tags)
	}
}
