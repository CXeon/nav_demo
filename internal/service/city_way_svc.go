package service

import (
	"github.com/CXeon/nav_demo/internal/entity"
	"github.com/CXeon/nav_demo/internal/model"
	"github.com/twpayne/go-geom"
)

var CityWaySvc *cityWayService

type cityWayService struct {
	*Service
}

func newCityWayService(svc *Service) {
	CityWaySvc = &cityWayService{
		svc,
	}
}

// CreateOne 创建一条路线数据
func (svc *cityWayService) CreateOne(data entity.CityWay) (id uint, err error) {

	lineStr := geom.NewLineString(geom.XY)
	lineStr.SetSRID(4326)

	coords := make([]geom.Coord, len(data.LonLat))

	for i, c := range data.LonLat {
		coords[i] = geom.Coord{c[0], c[1]}
	}

	lineString, err := lineStr.SetCoords(coords)
	if err != nil {
		return 0, err
	}

	//实体转模型
	m := model.CityWay{
		Geom:        lineString,
		Cost:        data.Cost,
		ReverseCost: data.ReverseCost,
		Source:      0,
		Target:      0,
	}
	id, err = svc.pgsqlDao.CityWayDao.InsertOne(m)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindOne 查找一条数据
func (svc *cityWayService) FindOne(id uint) (ent *entity.CityWay, err error) {
	one, err := svc.pgsqlDao.CityWayDao.FindOne(id)
	if err != nil {
		return nil, err
	}

	//lineString转数组
	lineString := one.Geom
	lonLat := make([][]float64, lineString.NumCoords())

	for i := 0; i < lineString.NumCoords(); i++ {
		lonLat[i] = make([]float64, 2)
		lonLat[i][0] = lineString.Coord(i)[0]
		lonLat[i][1] = lineString.Coord(i)[1]
	}

	ent = &entity.CityWay{
		ID:          one.ID,
		CreatedAt:   one.CreatedAt,
		UpdatedAt:   one.UpdatedAt,
		LonLat:      lonLat,
		Cost:        one.Cost,
		ReverseCost: one.ReverseCost,
		CreateAtStr: one.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAtStr: one.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ent, nil
}
