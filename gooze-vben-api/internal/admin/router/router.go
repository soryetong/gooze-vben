package router

import (
	"gooze-vben-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soryetong/gooze-starter/services/gzmiddleware"
	"github.com/spf13/viper"
)

func InitRouter() *gin.Engine {
	setMode()

	r := gin.New()
	fs := "/static"
	r.StaticFS(fs, http.Dir("./"+fs))

	r.Use(gzmiddleware.Begin()).Use(gzmiddleware.ErrorHandler()).Use(gzmiddleware.Cross()).Use(middleware.Record())
	publicGroup := r.Group("api/v1")
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})

		InitSystemAuthPublicRouter(publicGroup)
	}

	privateAuthGroup := r.Group("api/v1")
	privateAuthGroup.Use(gzmiddleware.Jwt()).Use(gzmiddleware.Casbin())
	{
		InitMaterialAuthRouter(privateAuthGroup)
		InitApiAuthRouter(privateAuthGroup)
		InitDictAuthRouter(privateAuthGroup)
		InitMenuAuthRouter(privateAuthGroup)
		InitRecordAuthRouter(privateAuthGroup)
		InitRoleAuthRouter(privateAuthGroup)
		InitUserAuthRouter(privateAuthGroup)
	}

	return r
}

func setMode() {
	switch viper.GetString("App.Env") {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.TestMode)
	}
}
