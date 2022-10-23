package api

import (
	"MyToDoList/pkg/util"
	"MyToDoList/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var us *service.UserService
	if err := c.ShouldBindJSON(us); err != nil {
		util.LogrusObj.Info(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := us.Register()
		c.JSON(200, res)
	}
}

func UserLogin(c *gin.Context) {
	var us *service.UserService
	if err := c.ShouldBindJSON(us); err != nil {
		util.LogrusObj.Info(err)
		c.JSON(400, ErrorResponse(err))
	} else {
		res := us.Login()
		c.JSON(200, res)
	}
}
