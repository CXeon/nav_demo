package http

import (
	"github.com/CXeon/nav_demo/config"
	"github.com/gin-gonic/gin"
)

func initRoutes(conf *config.Config) *gin.Engine {
	r := gin.Default()

	return r
}
