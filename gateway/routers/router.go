package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/fy026/weather/gateway/middleware/jwt"
	"github.com/fy026/weather/gateway/pkg/setting"
	"github.com/fy026/weather/gateway/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/test", api.GetTest)
	}

	return r
}
