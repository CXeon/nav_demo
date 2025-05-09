package dao

import "github.com/CXeon/micro_contrib/postgresql"

type CityWayDao struct {
	*postgresql.Client
}

func newCityWayDao(cli *postgresql.Client) *CityWayDao {
	return &CityWayDao{
		cli,
	}
}
