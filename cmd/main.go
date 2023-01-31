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
	var cnt int = 0

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
			if e.Tags.HasTag("railway") {
				cnt = cnt + 1
				fmt.Println(e.Tags)
			}
		case *osm.Way:
			if e.Tags.HasTag("railway") {
				cnt = cnt + 1
				fmt.Println(e.Tags)
			}
		case *osm.Relation:
			if e.Tags.HasTag("railway") {
				cnt = cnt + 1
				fmt.Println(e.Tags)
			}
		}
	}
	end := time.Now()
	fmt.Println(end)
	fmt.Println(time.Since(begin))
	fmt.Println("ending")
	fmt.Println("theres ", cnt, "objects that have rails in it")
	fmt.Println("scanning ", i, "objects from your file, thank me bitch")

	switch os.Args[1] {
	case "read":
	case "compare":
	}

}
