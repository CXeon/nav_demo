package dao

import "github.com/CXeon/micro_contrib/postgresql"

type PostgresqlCloudDao struct {
	client     *postgresql.Client
	CityWayDao *CityWayDao
}

func NewPostgresqlCloudDao(client *postgresql.Client) *PostgresqlCloudDao {
	return &PostgresqlCloudDao{
		client:     client,
		CityWayDao: newCityWayDao(client),
	}
}
