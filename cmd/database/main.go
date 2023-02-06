package main

import (
	"log"
	"railsearch/pkg/config"
	"railsearch/pkg/database"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230205 = gormigrate.Migration{
	ID: "V20230205",
	Migrate: func(tx *gorm.DB) error {
		type Node struct {
			NodeId    int64 `gorm:"primaryKey"`
			Latitude  float64
			Longitude float64
		}

		type IndexNode struct {
			NodeId     int64   `gorm:"primaryKey"`
			CartesianX float64 `gorm:"index:cartesian_way"`
			CartesianY float64 `gorm:"index:cartesian_way"`
		}

		type WayNode struct {
			ID     uint  `gorm:"primaryKey"`
			WayId  int64 `gorm:"index"`
			NodeId int64 `gorm:"index"`
			Order  int
		}

		type TargetNode struct {
			OsmId     int64 `gorm:"primaryKey"`
			Latitude  float64
			Longitude float64
			Tags      string
			Type      string `gorm:"type:varchar(10)"`
		}

		return tx.AutoMigrate(&Node{}, &IndexNode{}, &WayNode{}, TargetNode{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("nodes", "target_nodes", "way_nodes")
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
