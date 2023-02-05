package main

import (
	"context"
	"fmt"

	"github.com/paulmach/osm/osmapi"
	"golang.org/x/time/rate"
)

func main() {
	ctx := context.Background()

	osmapi.DefaultDatasource.Limiter = rate.NewLimiter(10, 1)

	fun, _ := osmapi.Node(ctx, 10536442398)
	fmt.Println("-------------")
	fmt.Println("id ", fun.ID)
	fmt.Println("lat ", fun.Lat)
	fmt.Println("lon ", fun.Lon)
	fmt.Println("user ", fun.User)
	fmt.Println("userid ", fun.UserID)
	fmt.Println("tags ", fun.Tags)
	fmt.Println("-------------")
}
