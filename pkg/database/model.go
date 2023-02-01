package database

import "gorm.io/gorm"

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
