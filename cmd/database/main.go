package main

import (
	"log"
	"railsearch/pkg/config"
	"railsearch/pkg/database"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230205 = gormigrate.Migration{
	ID: "V20230201",
	Migrate: func(tx *gorm.DB) error {
		type Node struct {
			gorm.Model
			NodeId     int64   `gorm:"index"`
			Latitude   float64 `gorm:"index:location_node"`
			Longitude  float64 `gorm:"index:location_node"`
			CartesianX float64 `gorm:"index:cartesian_node"`
			CartesianY float64 `gorm:"index:cartesian_node"`
			IsBuilding bool    `gorm:"index"`
		}

		type WayMember struct {
			gorm.Model
			WayId  int64 `gorm:"index"`
			NodeId int64 `gorm:"index"`
			Order  int
		}

		type BuildingNode struct {
			gorm.Model
			OsmId     int64   `gorm:"index"`
			Latitude  float64 `gorm:"index:location_way"`
			Longitude float64 `gorm:"index:location_way"`
			OsmType   string  `gorm:"index"`
			Tags      string  `gorm:"index"`
		}

		return tx.AutoMigrate(&Node{}, &WayMember{}, &BuildingNode{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("nodes", "way_members", "building_nodes")
	},
}

func main() {
	config := config.GetConfig("config.json")
	database.InitDB(config.DatabaseConfig)

	m := gormigrate.New(database.GetConn(), gormigrate.DefaultOptions, []*gormigrate.Migration{
		&V20230205,
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
