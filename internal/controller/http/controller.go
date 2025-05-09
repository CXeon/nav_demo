package http

import (
	"github.com/CXeon/nav_demo/config"
)

type controller struct {
	conf *config.Config
}

func NewController(conf *config.Config) *controller {
	if conf == nil {
		return nil
	}

	return &controller{
		conf: conf,
	}
}
