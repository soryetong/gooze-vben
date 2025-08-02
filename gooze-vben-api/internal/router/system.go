package router

import (
	"github.com/gin-gonic/gin"
	"gooze-vben-api/internal/handler"
)

func InitApiAuthRouter(routerGroup *gin.RouterGroup) {
	apiGroup := routerGroup.Group("/api")
	{
		apiGroup.POST("/add", handler.ApiAdd)
		apiGroup.GET("/list", handler.ApiList)
		apiGroup.PUT("/update/:id", handler.ApiUpdate)
		apiGroup.DELETE("/delete/:id", handler.ApiDelete)
	}
}

func InitDictAuthRouter(routerGroup *gin.RouterGroup) {
	dictGroup := routerGroup.Group("/dict")
	{
		dictGroup.POST("/add", handler.DictAdd)
		dictGroup.GET("/list", handler.DictList)
		dictGroup.PUT("/update/:id", handler.DictUpdate)
		dictGroup.DELETE("/delete/:id", handler.DictDelete)
	}
}

func InitMenuAuthRouter(routerGroup *gin.RouterGroup) {
	menuGroup := routerGroup.Group("/menu")
	{
		menuGroup.GET("/router", handler.MenuRouter)
		menuGroup.GET("/tree", handler.MenuTree)
		menuGroup.POST("/add", handler.MenuAdd)
		menuGroup.PUT("/update/:id", handler.MenuUpdate)
		menuGroup.GET("/info/:id", handler.MenuInfo)
		menuGroup.DELETE("/delete/:id", handler.MenuDelete)
	}
}

func InitRecordAuthRouter(routerGroup *gin.RouterGroup) {
	recordGroup := routerGroup.Group("/record")
	{
		recordGroup.GET("/list", handler.RecordList)
		recordGroup.DELETE("/delete/:id", handler.RecordDelete)
	}
}

func InitRoleAuthRouter(routerGroup *gin.RouterGroup) {
	roleGroup := routerGroup.Group("/role")
	{
		roleGroup.POST("/add", handler.RoleAdd)
		roleGroup.GET("/list", handler.RoleList)
		roleGroup.GET("/info/:id", handler.RoleInfo)
		roleGroup.PUT("/update/:id", handler.RoleUpdate)
		roleGroup.PUT("/assign/:id", handler.RoleAssign)
		roleGroup.DELETE("/delete/:id", handler.RoleDelete)
	}
}

func InitUserAuthRouter(routerGroup *gin.RouterGroup) {
	userGroup := routerGroup.Group("/user")
	{
		userGroup.GET("/info", handler.UserInfo)
		userGroup.POST("/add", handler.UserAdd)
		userGroup.GET("/list", handler.UserList)
		userGroup.PUT("/update/:id", handler.UserUpdate)
		userGroup.DELETE("/delete/:id", handler.UserDelete)
	}
}
