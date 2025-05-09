package service

import (
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/nav_demo/config"
	"github.com/CXeon/nav_demo/internal/dao"
)

type Service struct {
	conf     *config.Config
	logger   *log.Logger
	pgsqlDao *dao.PostgresqlCloudDao
}

func NewService(conf *config.Config, logger *log.Logger, pgsqlDao *dao.PostgresqlCloudDao) error {
	svc := &Service{
		conf:     conf,
		logger:   logger,
		pgsqlDao: pgsqlDao,
	}

	newCityWayService(svc)

	return nil
}
