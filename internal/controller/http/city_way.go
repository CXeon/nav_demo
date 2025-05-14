package http

import (
	"github.com/CXeon/micro_contrib/err_code"
	"github.com/CXeon/micro_contrib/errors"
	"github.com/CXeon/micro_contrib/response"
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
		response.ResponseError(ctx, http.StatusBadRequest, errors.Wrap(err_code.ErrCommonServer, err.Error()))
		return
	}

	//执行业务
	id, err := service.CityWaySvc.CreateOne(data)
	if err != nil {
		response.ResponseError(ctx, http.StatusInternalServerError, errors.Wrap(err_code.ErrCommonServer, err.Error()))
		return
	}
	response.ResponseSuccess(ctx, map[string]uint{"id": id})
	return
}

// FindOne 查找一条数据
func (c *controller) FindOne(ctx *gin.Context) {
	var err error

	//解析请求参数
	idStr := ctx.DefaultQuery("id", "")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, errors.Wrap(err_code.ErrCommonServer, err.Error()))
		return
	}

	//执行业务
	one, err := service.CityWaySvc.FindOne(uint(idInt))
	if err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, errors.Wrap(err_code.ErrCommonServer, err.Error()))
		return
	}
	response.ResponseSuccess(ctx, one)
	return
}
