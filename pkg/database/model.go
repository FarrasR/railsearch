package database

import "gorm.io/gorm"

//Latitude is the Y axis, longitude is the X axis

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

type WayMember struct {
	gorm.Model
	WayId  int64 `gorm:"index"`
	NodeId int64 `gorm:"index"`
	Order  int
}
