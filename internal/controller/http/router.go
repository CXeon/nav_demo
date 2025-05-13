package http

import (
	"fmt"
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/nav_demo/config"
	"github.com/gin-gonic/gin"
)

func initRoutes(conf *config.Config, logger *log.Logger) *gin.Engine {
	r := gin.Default()
	flmGroup := r.Group("/flm")
	appGroup := flmGroup.Group(fmt.Sprintf("/%s", conf.Application.ServiceName))

	//创建controller
	c := NewController(conf, logger)

	cityGroup := appGroup.Group("/city")
	{
		cityGroup.POST("/way", c.CreateOne)
		cityGroup.GET("/way", c.FindOne)
	}

	return r
}
