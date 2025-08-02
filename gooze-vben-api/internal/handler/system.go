package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/soryetong/gooze-starter/gooze"
	"github.com/soryetong/gooze-starter/pkg/gzerror"
	"github.com/soryetong/gooze-starter/pkg/gzutil"
	"github.com/spf13/cast"
	"gooze-vben-api/internal/dto"
	"gooze-vben-api/internal/logic"
)

var systemLogic = logic.NewSystemLogic()

// @Summary ApiAdd
// @Description ApiAdd
// @Accept json
// @Produce json
// @Param body dto.UpsertApiReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /api/add [post]
func ApiAdd(ctx *gin.Context) {
	var req dto.UpsertApiReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.ApiAdd(ctx, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary ApiList
// @Description ApiList
// @Accept json
// @Produce json
// @Param query dto.ApiListReq
// @Success 200 {object} dto.ApiListResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /api/list [get]
func ApiList(ctx *gin.Context) {
	var req dto.ApiListReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	resp, err := systemLogic.ApiList(ctx, &req)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary ApiUpdate
// @Description ApiUpdate
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Param body dto.UpsertApiReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /api/update/:id [put]
func ApiUpdate(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	var req dto.UpsertApiReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.ApiUpdate(ctx, id, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary ApiDelete
// @Description ApiDelete
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /api/delete/:id [delete]
func ApiDelete(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := systemLogic.ApiDelete(ctx, id); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary DictAdd
// @Description DictAdd
// @Accept json
// @Produce json
// @Param body dto.UpsertDictReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /dict/add [post]
func DictAdd(ctx *gin.Context) {
	var req dto.UpsertDictReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.DictAdd(ctx, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary DictList
// @Description DictList
// @Accept json
// @Produce json
// @Param query dto.DictListReq
// @Success 200 {object} dto.DictListResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /dict/list [get]
func DictList(ctx *gin.Context) {
	var req dto.DictListReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	resp, err := systemLogic.DictList(ctx, &req)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary DictUpdate
// @Description DictUpdate
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Param body dto.UpsertDictReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /dict/update/:id [put]
func DictUpdate(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	var req dto.UpsertDictReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.DictUpdate(ctx, id, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary DictDelete
// @Description DictDelete
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /dict/delete/:id [delete]
func DictDelete(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := systemLogic.DictDelete(ctx, id); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary MenuRouter
// @Description MenuRouter
// @Accept json
// @Produce json
// @Success 200 {object} dto.MenuResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /menu/router [get]
func MenuRouter(ctx *gin.Context) {
	resp, err := systemLogic.MenuRouter(ctx)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary MenuTree
// @Description MenuTree
// @Accept json
// @Produce json
// @Param query dto.MenuTreeReq
// @Success 200 {object} dto.MenuResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /menu/tree [get]
func MenuTree(ctx *gin.Context) {
	var req dto.MenuTreeReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	resp, err := systemLogic.MenuTree(ctx, &req)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary MenuAdd
// @Description MenuAdd
// @Accept json
// @Produce json
// @Param body dto.MenuInfo
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /menu/add [post]
func MenuAdd(ctx *gin.Context) {
	var req dto.MenuInfo
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.MenuAdd(ctx, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary MenuUpdate
// @Description MenuUpdate
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Param body dto.MenuInfo
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /menu/update/:id [put]
func MenuUpdate(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	var req dto.MenuInfo
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.MenuUpdate(ctx, id, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary MenuInfo
// @Description MenuInfo
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 {object} dto.MenuInfo
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /menu/info/:id [get]
func MenuInfo(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	resp, err := systemLogic.MenuInfo(ctx, id)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary MenuDelete
// @Description MenuDelete
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /menu/delete/:id [delete]
func MenuDelete(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := systemLogic.MenuDelete(ctx, id); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary RecordList
// @Description RecordList
// @Accept json
// @Produce json
// @Param query dto.RecordListReq
// @Success 200 {object} dto.RecordListResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /record/list [get]
func RecordList(ctx *gin.Context) {
	var req dto.RecordListReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	resp, err := systemLogic.RecordList(ctx, &req)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary RecordDelete
// @Description RecordDelete
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /record/delete/:id [delete]
func RecordDelete(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := systemLogic.RecordDelete(ctx, id); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary RoleAdd
// @Description RoleAdd
// @Accept json
// @Produce json
// @Param body dto.UpsertRoleReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /role/add [post]
func RoleAdd(ctx *gin.Context) {
	var req dto.UpsertRoleReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.RoleAdd(ctx, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary RoleList
// @Description RoleList
// @Accept json
// @Produce json
// @Param query dto.RoleListReq
// @Success 200 {object} dto.RoleListResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /role/list [get]
func RoleList(ctx *gin.Context) {
	var req dto.RoleListReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	resp, err := systemLogic.RoleList(ctx, &req)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary RoleInfo
// @Description RoleInfo
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 {object} dto.RoleInfoResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /role/info/:id [get]
func RoleInfo(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	resp, err := systemLogic.RoleInfo(ctx, id)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary RoleUpdate
// @Description RoleUpdate
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Param body dto.UpsertRoleReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /role/update/:id [put]
func RoleUpdate(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	var req dto.UpsertRoleReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.RoleUpdate(ctx, id, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary RoleAssign
// @Description RoleAssign
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Param body dto.AssignRoleReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /role/assign/:id [put]
func RoleAssign(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	var req dto.AssignRoleReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.RoleAssign(ctx, id, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary RoleDelete
// @Description RoleDelete
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /role/delete/:id [delete]
func RoleDelete(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := systemLogic.RoleDelete(ctx, id); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary UserInfo
// @Description UserInfo
// @Accept json
// @Produce json
// @Success 200 {object} dto.UserInfoResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /user/info [get]
func UserInfo(ctx *gin.Context) {
	resp, err := systemLogic.UserInfo(ctx)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary UserAdd
// @Description UserAdd
// @Accept json
// @Produce json
// @Param body dto.UpsertUserReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /user/add [post]
func UserAdd(ctx *gin.Context) {
	var req dto.UpsertUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.UserAdd(ctx, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary UserList
// @Description UserList
// @Accept json
// @Produce json
// @Param query dto.UserListReq
// @Success 200 {object} dto.UserListResp
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /user/list [get]
func UserList(ctx *gin.Context) {
	var req dto.UserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	resp, err := systemLogic.UserList(ctx, &req)
	if err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, resp)
}

// @Summary UserUpdate
// @Description UserUpdate
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Param body dto.UpsertUserReq
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /user/update/:id [put]
func UserUpdate(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	var req dto.UpsertUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		gooze.FailWithMessage(ctx, gzerror.Trans(err))
		return
	}

	if err := systemLogic.UserUpdate(ctx, id, &req); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}

// @Summary UserDelete
// @Description UserDelete
// @Accept json
// @Produce json
// @Param id query int64 true "id"
// @Success 200 string success
// @Failure 200 {object} gooze.Response 根据Code表示不同类型的错误
// @Router /user/delete/:id [delete]
func UserDelete(ctx *gin.Context) {
	id := cast.ToInt64(ctx.Param("id"))
	if !gzutil.IsValidNumber(id) {
		gooze.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := systemLogic.UserDelete(ctx, id); err != nil {
		gooze.FailWithMessage(ctx, err.Error())
		return
	}
	gooze.Success(ctx, nil)
}
