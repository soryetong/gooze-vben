package logic

import (
	"context"
	"fmt"
	"gooze-vben-api/internal/admin/dto"
	"gooze-vben-api/internal/common"
	"gooze-vben-api/models"
	"path/filepath"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/soryetong/gooze-starter/gooze"
	"github.com/soryetong/gooze-starter/services/gzdb"
	"github.com/soryetong/gooze-starter/pkg/gzutil"
	"github.com/spf13/viper"
)

type MaterialLogic struct {
}

func NewMaterialLogic() *MaterialLogic {
	return &MaterialLogic{}
}

// @Summary MaterialAdd
func (self *MaterialLogic) MaterialAdd(ctx context.Context, params *dto.MaterialUpsertReq) (err error) {
	fileExt := filepath.Ext(params.Resource.Filename)
	if err = common.CheckIAVExt(ctx, models.MaterialType(params.Type), fileExt); err != nil {
		return err
	}

	fileUrl, err := common.IAVSave(ctx, params.Resource, "material")
	if err != nil {
		return err
	}

	data := new(models.CMaterials)
	_ = copier.Copy(data, params)
	data.ResourceUrl = strings.TrimLeft(fileUrl, ".")
	err = gooze.Gorm().Model(&models.CMaterials{}).Create(data).Error

	return
}

// @Summary MaterialList
func (self *MaterialLogic) MaterialList(ctx context.Context, params *dto.MaterialListReq) (resp *dto.CommonListResp, err error) {
	resp = &dto.CommonListResp{}

	query := gooze.Gorm().Model(&models.CMaterials{}).Order("id desc")
	query.Where("type = ?", params.Type)
	if params.Status > 0 {
		query.Where("status = ?", params.Status)
	}
	if params.Keyword != "" {
		query.Where("name like ?", "%"+params.Keyword+"%")
	}
	if err = query.Count(&resp.Total).Error; err != nil {
		return
	}

	var list []*models.CMaterials
	if err = query.Scopes(gzdb.GormPaginate(params.Page, params.PageSize)).Find(&list).Error; err != nil {
		return
	}

	items := make([]*dto.MaterialInfoResp, 0)
	for _, info := range list {
		item := new(dto.MaterialInfoResp)
		_ = copier.Copy(item, info)
		item.Resource = gzutil.JoinDomain(viper.GetString("oss.url"), info.ResourceUrl)
		item.CreateTime = info.CreatedAt.UnixMilli()
		items = append(items, item)
	}
	resp.Items = items

	return
}

// @Summary MaterialUpdate
func (self *MaterialLogic) MaterialUpdate(ctx context.Context, id int64, params *dto.MaterialUpsertReq) (err error) {
	var info *models.CMaterials
	if err = gooze.Gorm().Model(&models.CMaterials{}).Where("id = ?", id).Find(&info).Error; err != nil {
		return fmt.Errorf("资源不存在")
	}

	update := map[string]interface{}{
		"name":        params.Name,
		"description": params.Description,
		"status":      params.Status,
	}
	if params.Resource != nil && params.Resource.Size > 0 {
		fileExt := filepath.Ext(params.Resource.Filename)
		if err = common.CheckIAVExt(ctx, models.MaterialType(params.Type), fileExt); err != nil {
			return err
		}

		fileUrl, err := common.IAVSave(ctx, params.Resource, "material")
		if err != nil {
			return err
		}
		update["resource_url"] = strings.TrimLeft(fileUrl, ".")
	}

	err = gooze.Gorm().Model(&models.CMaterials{}).Where("id = ?", id).Updates(update).Error

	return
}

// @Summary MaterialDelete
func (self *MaterialLogic) MaterialDelete(ctx context.Context, id int64) (err error) {
	err = gooze.Gorm().Delete(&models.CMaterials{}, "id = ?", id).Error

	return
}
