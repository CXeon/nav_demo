package http

import (
	"github.com/CXeon/nav_demo/internal/entity"
	"github.com/CXeon/nav_demo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *controller) CreateOne(ctx *gin.Context) {
	var err error

	//解析请求参数
	data := entity.CityWay{}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//执行业务
	id, err := service.CityWaySvc.CreateOne(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
	return
}

// FindOne 查找一条数据
func (c *controller) FindOne(ctx *gin.Context) {
	var err error

	//解析请求参数
	idStr := ctx.DefaultQuery("id", "")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//执行业务
	one, err := service.CityWaySvc.FindOne(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": one})
	return
}
