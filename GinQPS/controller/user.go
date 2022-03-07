package controller

import (
	"GinQPS/config"
	"GinQPS/model"
	"GinQPS/service"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func AddGroupUser(r *gin.Engine) {
	group := r.Group("/user")
	group.GET("/get/id/:id", wrapper(GetOneUserById))
	group.GET("/get/id/redis/:id", wrapper(GetOneUserByIdOnRedis))
	group.GET("/get/name/:name", wrapper(GetOneUserByName))
	group.GET("/search/favorite/:favorite", wrapper(GetUserListBySearchFavorite))
	group.POST("/create", wrapper(CreateUser))
	group.GET("/list", wrapper(ListUsers))
}

func GetOneUserById(c *gin.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return EmptyResponse(c, config.REQUEST_ERR_CODE, err.Error())
	}
	user := service.GetUserById(id)
	return Response(c, config.SUCCESS_CODE, "success", user)
}

func GetOneUserByIdOnRedis(c *gin.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return EmptyResponse(c, config.REQUEST_ERR_CODE, err.Error())
	}
	user := service.GetUserByIdOnRedis(id)
	return Response(c, config.SUCCESS_CODE, "success", user)
}

func GetOneUserByName(c *gin.Context) error {
	name := c.Param("name")
	user := service.GetUserByName(name)
	return Response(c, config.SUCCESS_CODE, "success", user)
}

func GetUserListBySearchFavorite(c *gin.Context) error {
	favorite := c.Param("favorite")
	var users []model.User
	err := hystrix.Do("default_command", func() error {
		users = service.GetUserListBySearchFavorite(favorite)
		return nil
	}, func(err error) error {
		// 这里要处理，当超时或者返回的error不为空后，执行该逻辑。注意超时后，原本的逻辑仍然执行（比如sql，单独的goroutine），所以原来的逻辑那里不能返回请求。
		log.Printf("error occur: %s\n", err)
		c.Status(500)
		return err
	})
	if err != nil {
		return EmptyResponse(c, config.SERVER_ERR_CODE, err.Error())
	}
	return Response(c, config.SUCCESS_CODE, "success", len(users))
}

func CreateUser(c *gin.Context) error {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return EmptyResponse(c, config.REQUEST_ERR_CODE, "param error")
	}
	service.InsertUser(user)
	return Response(c, config.SUCCESS_CODE, "success", user)
}

func ListUsers(c *gin.Context) error {
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	name := c.Query("name")
	users := service.ListUsers(int(offset), int(limit), name)
	return Response(c, config.SUCCESS_CODE, "success", users)
}