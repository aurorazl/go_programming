package controller

import (
	"GinQPS/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			return
		}
	}
}

func Response(c *gin.Context, code int, msg string, data interface{}) error {
	res := model.CommonResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, res)
	return nil
}

func EmptyResponse(c *gin.Context, code int, msg string) error {
	res := model.CommonResult{
		Code: code,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, res)
	return nil
}
