package controller

import (
	"GinQPS/config"
	"GinQPS/model"
	"GinQPS/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddGroupUser(r *gin.Engine) {
	group := r.Group("/user")
	group.GET("/get/:id", wrapper(GetOneUser))
	group.POST("/create", wrapper(CreateUser))
}

func GetOneUser(c *gin.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return EmptyResponse(c, config.REQUEST_ERR_CODE, err.Error())
	}
	user := service.GetUserById(id)
	return Response(c, config.SUCCESS_CODE, "success", user)
}

func CreateUser(c *gin.Context) error {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return EmptyResponse(c, config.REQUEST_ERR_CODE, "param error")
	}
	service.InsertUser(user)
	return Response(c, config.SUCCESS_CODE, "success", user)
}
