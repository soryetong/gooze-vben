package logic

import (
	"context"
	"errors"
	"gooze-vben-api/internal/dto"
	"gooze-vben-api/models"
	"time"

	"github.com/soryetong/gooze-starter/gooze"
	"github.com/soryetong/gooze-starter/pkg/gzauth"
	"github.com/soryetong/gooze-starter/pkg/gzutil"
	"go.uber.org/zap"
)

type AuthLogic struct {
}

func NewAuthLogic() *AuthLogic {
	return &AuthLogic{}
}

// @Summary SystemAuthLogin
func (self *AuthLogic) SystemAuthLogin(ctx context.Context, params *dto.LoginReq) (resp *dto.LoginResp, err error) {
	user := new(models.SysUsers)
	if err = gooze.Gorm().Model(&models.SysUsers{}).Where("username = ?", params.Username).First(user).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	if user.Status != models.SysUserStatusNormal {
		return nil, errors.New("用户已被禁用")
	}

	if gzutil.ValidatePasswd(params.Password, user.Salt, user.Password) == false {
		return nil, errors.New("密码错误")
	}

	token, err := gzauth.GenerateJwtToken(map[string]interface{}{
		"id":       user.Id,
		"username": user.Username,
		"role_id":  user.RoleId,
		"exp":      time.Now().Add(time.Minute * 20).Unix(),
	})
	if err != nil {
		gooze.Log.Error("生成 token 失败", zap.Error(err))
		return nil, errors.New("生成 token 失败")
	}

	resp = &dto.LoginResp{
		AccessToken: token,
		Id:          user.Id,
		Password:    "",
		RealName:    user.Username,
		Roles:       []string{""},
		Username:    user.Username,
	}

	return
}

// @Summary SystemAuthLogout
func (self *AuthLogic) SystemAuthLogout(ctx context.Context, params *dto.LogoutReq) (err error) {
	// TODO implement

	return
}
