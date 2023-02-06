package database

//Latitude is the Y axis, longitude is the X axis

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
