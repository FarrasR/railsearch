package main

type CartesianCoordinate struct {
	CartesianX float64
	CartesianY float64
}

// var shift float64

// a blunder force me to write this

//orb.Point{longitude, latitude}
//Latitude is the Y axis, longitude is the X axis
func main() {
	// config := config.GetConfig("config.json")
	// database.InitDB(config.DatabaseConfig)

	// rows, err := database.GetConn().Model(&database.Way{}).Rows()
	// if err != nil {
	// 	panic(err)
	// }

	// i := 0
	// for rows.Next() {
	// 	i = i + 1

	// 	if i%100 == 0 {
	// 		fmt.Println("updating ", i, " ways")
	// 	}
	// 	var way database.Way
	// 	database.GetConn().ScanRows(rows, &way)

	// 	mercator := project.Point(orb.Point{way.Longitude, way.Latitude}, project.WGS84.ToMercator)
	// 	way.CartesianX = mercator[0]
	// 	way.CartesianY = mercator[1]
	// 	database.GetConn().Save(way)
	// }

	// sf := orb.Point{110.41407570000001, -6.973148}
	// merc := project.Point(sf, project.WGS84.ToMercator)
	// fmt.Println("geo 1", sf)
	// fmt.Println("merc 1", merc)

	// sf2 := orb.Point{110.41514910000001, -6.9723426}
	// merc2 := project.Point(sf2, project.WGS84.ToMercator)
	// fmt.Println("geo 2", sf)
	// fmt.Println("merc 2", merc2)

	// distmerc := math.Sqrt(math.Abs(merc[0]-merc2[0])*math.Abs(merc[0]-merc2[0]) + math.Abs(merc[1]-merc2[1])*math.Abs(merc[1]-merc2[1]))
	// fmt.Println("distmerc ", distmerc)

	// distgeo := geo.Distance(sf, sf2)
	// fmt.Println("distgeo ", distgeo)
}

// func Converter(longitude float64, latitude float64) CartesianCoordinate {
// 	fmt.Println(longitude * shift / 180)
// 	return CartesianCoordinate{
// 		CartesianX: longitude * shift / 180,
// 		CartesianY: math.Log(math.Tan((90+latitude)*math.Pi/360)) / (math.Pi / 180),
// 	}
// }
