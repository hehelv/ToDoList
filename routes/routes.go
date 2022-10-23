package routes

import (
	"MyToDoList/api"
	"MyToDoList/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // 默认是debug模式
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

	}
	return r
}
