package main

import (
	"log"
	"railsearch/pkg/config"
	"railsearch/pkg/database"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230201 = gormigrate.Migration{
	ID: "V20230201",
	Migrate: func(tx *gorm.DB) error {
		type Node struct {
			gorm.Model
			NodeId    int64   `gorm:"index"`
			Latitude  float64 `gorm:"index:location_node"`
			Longitude float64 `gorm:"index:location_node"`
		}

		type Way struct {
			gorm.Model
			WayId     int64   `gorm:"index"`
			Latitude  float64 `gorm:"index:location_way"`
			Longitude float64 `gorm:"index:location_way"`
		}

		type WayMember struct {
			gorm.Model
			WayId  int64 `gorm:"index"`
			NodeId int64 `gorm:"index"`
			Order  int
		}

		return tx.AutoMigrate(&Node{}, &Way{}, &WayMember{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("node", "way", "way_member")
	},
}

func main() {
	config := config.GetConfig("config.json")
	database.InitDB(config.DatabaseConfig)

	m := gormigrate.New(database.GetConn(), gormigrate.DefaultOptions, []*gormigrate.Migration{
		&V20230201,
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
