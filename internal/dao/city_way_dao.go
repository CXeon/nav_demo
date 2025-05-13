package dao

import (
	"github.com/CXeon/micro_contrib/postgresql"
	"github.com/CXeon/nav_demo/internal/model"
)

type CityWayDao struct {
	*postgresql.Client
}

func newCityWayDao(cli *postgresql.Client) *CityWayDao {
	return &CityWayDao{
		cli,
	}
}

// InsertOne 插入一条记录
func (dao *CityWayDao) InsertOne(cityWay model.CityWay) (id uint, err error) {

	db := dao.Db

	if result := db.Debug().Create(&cityWay); result.Error != nil {
		return 0, result.Error
	}
	return cityWay.ID, nil
}

// FindOne 查找一条记录
func (dao *CityWayDao) FindOne(id uint) (*model.CityWay, error) {
	db := dao.Db
	m := model.CityWay{}
	result := db.Select([]string{"id", "geom", "cost", "reverse_cost", "source", "target", "created_at", "updated_at"}).
		Where("deleted_at is NULL").First(&m, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &m, nil
}
