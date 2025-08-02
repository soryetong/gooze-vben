package router

import (
	"github.com/gin-gonic/gin"
	"gooze-vben-api/internal/handler"
)

func InitSystemAuthPublicRouter(routerGroup *gin.RouterGroup) {
	systemAuthGroup := routerGroup.Group("/system/auth")
	{
		systemAuthGroup.POST("/login", handler.SystemAuthLogin)
		systemAuthGroup.POST("/logout", handler.SystemAuthLogout)
	}
}
