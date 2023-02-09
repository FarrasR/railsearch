package database

//Latitude is the Y axis, longitude is the X axis

// node lists ALL nodes in the file including those that isnt needed
type Node struct {
	NodeId    int64 `gorm:"primaryKey"`
	Latitude  float64
	Longitude float64
}

//indexNode lists all node that contains the search tag
type IndexNode struct {
	NodeId     int64   `gorm:"primaryKey"`
	CartesianX float64 `gorm:"index:cartesian_way"`
	CartesianY float64 `gorm:"index:cartesian_way"`
}

//wayNode is lists of node that is used by ways that has the search tag
type WayNode struct {
	ID     uint  `gorm:"primaryKey"`
	WayId  int64 `gorm:"index"`
	NodeId int64 `gorm:"index"`
	Order  int
}

//targetNode is the list of node that around the node
type TargetNode struct {
	OsmId     int64 `gorm:"primaryKey"`
	Latitude  float64
	Longitude float64
	Tags      string
	Type      string `gorm:"type:varchar(10)"`
}
