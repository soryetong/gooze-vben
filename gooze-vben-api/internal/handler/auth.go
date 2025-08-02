package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/soryetong/gooze-starter/gooze"
	"github.com/soryetong/gooze-starter/pkg/gzerror"
	"gooze-vben-api/internal/dto"
	"gooze-vben-api/internal/logic"
)

var authLogic = logic.NewAuthLogic()

// @Summary SystemAuthLogin
// @Description SystemAuthLogin
// @Accept json
// @Produce json
// @Param body dto.LoginReq
// @Success 200 {object} dto.LoginResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /system/auth/login [post]
func SystemAuthLogin(ctx *gin.Context) {
	var req dto.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	resp, err := authLogic.SystemAuthLogin(ctx, &req)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary SystemAuthLogout
// @Description SystemAuthLogout
// @Accept json
// @Produce json
// @Param body dto.LogoutReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /system/auth/logout [post]
func SystemAuthLogout(ctx *gin.Context) {
	var req dto.LogoutReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := authLogic.SystemAuthLogout(ctx, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}
