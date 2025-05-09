package model

import (
	"github.com/twpayne/go-geom"
	"gorm.io/gorm"
)

type CityWay struct {
	gorm.Model
	Geom        geom.LineString
	Cost        float64
	ReverseCost float64
	Source      int
	Target      int
}

func (m *CityWay) TableName() string {
	return "tb_city_way"
}
