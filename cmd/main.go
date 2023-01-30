package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
)

func main() {
	f, err := os.Open("./java-latest.osm.pbf")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()

	// for scanner.Scan() {
	// 	o := scanner.Object()
	// 	o.(type)
	// 	// do something
	// }
	var i int = 0

	begin := time.Now()
	fmt.Println("beginning scan")
	fmt.Println(begin)
	for scanner.Scan() {
		i = i + 1
		if i%1000000 == 0 {
			fmt.Println("searching... count ", i)
		}

		switch e := scanner.Object().(type) {
		case *osm.Node:
			// i = i + 1
			// fmt.Println("-------------------")
			// fmt.Println(e.ID)
			// fmt.Println(e.XMLName.Local)
			// fmt.Println(e.XMLName.Space)
			// fmt.Println(e.Tags)
			// fmt.Println("-------------------")

			if e.Tags.HasTag("train") {
				fmt.Println(e)
			}

		case *osm.Way:
		case *osm.Relation:
		}
	}
	end := time.Now()
	fmt.Println(end)
	fmt.Println(time.Since(begin))
	fmt.Println("ending")

	fmt.Println("scanning ", i, "objects from your file, thank me bitch")
}
