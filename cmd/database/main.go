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

var V20230203 = gormigrate.Migration{
	ID: "V20230203",
	Migrate: func(tx *gorm.DB) error {
		type Node struct {
			gorm.Model
			NodeId     int64   `gorm:"index"`
			Latitude   float64 `gorm:"index:location_node"`
			Longitude  float64 `gorm:"index:location_node"`
			CartesianX float64 `gorm:"index:cartesian_node"`
			CartesianY float64 `gorm:"index:cartesian_node"`
		}

		type Way struct {
			gorm.Model
			WayId      int64   `gorm:"index"`
			Latitude   float64 `gorm:"index:location_way"`
			Longitude  float64 `gorm:"index:location_way"`
			CartesianX float64 `gorm:"index:cartesian_way"`
			CartesianY float64 `gorm:"index:cartesian_way"`
		}

		return tx.AutoMigrate(&Node{}, &Way{})
	},
	Rollback: func(tx *gorm.DB) error {
		err := tx.Migrator().DropIndex("nodes", "cartesian_node")
		if err != nil {
			return err
		}

		err = tx.Migrator().DropIndex("ways", "cartesian_way")
		if err != nil {
			return err
		}

		err = tx.Migrator().DropColumn("nodes", "cartesian_x")
		if err != nil {
			return err
		}

		err = tx.Migrator().DropColumn("nodes", "cartesian_y")
		if err != nil {
			return err
		}

		err = tx.Migrator().DropColumn("ways", "cartesian_x")
		if err != nil {
			return err
		}

		err = tx.Migrator().DropColumn("ways", "cartesian_y")
		if err != nil {
			return err
		}

		return nil
	},
}

func main() {
	config := config.GetConfig("config.json")
	database.InitDB(config.DatabaseConfig)

	m := gormigrate.New(database.GetConn(), gormigrate.DefaultOptions, []*gormigrate.Migration{
		&V20230201,
		&V20230203,
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
