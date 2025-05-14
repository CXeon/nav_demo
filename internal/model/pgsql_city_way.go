package model

import (
	"github.com/CXeon/micro_contrib/gorm/datatypes/geometry"
	"gorm.io/gorm"
)

type CityWay struct {
	gorm.Model
	Geom        *geometry.GeoLine `gorm:"column:geom;type:geometry(LineString, 4326);not null;index;comment:路线;"`
	Cost        float64           `gorm:"column:cost;type:float8;not null;comment:路线方向"`
	ReverseCost float64           `gorm:"column:reverse_cost;type:float8;not null;comment:路线反方向"`
	Source      int               `gorm:"column:source;type:int8;not null;comment:路网source"`
	Target      int               `gorm:"column:target;type:int8;not null;comment:路网target"`
}

func (m *CityWay) TableName() string {
	return tableNamePrefix() + "city_way"
}
