package entity

import (
	"time"
)

type CityWay struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//Geom        geom.LineString `json:"geom"`
	LonLat      [][]float64 `json:"lon_lat"` //路线各个点的经纬度
	Cost        float64     `json:"cost"`
	ReverseCost float64     `json:"reverse_cost"`
	//Source      int             `json:"source"`
	//Target      int             `json:"target"`
	CreateAtStr string `json:"create_at_str"`
	UpdateAtStr string `json:"update_at_str"`
}
