package common

import (
	"context"
	"errors"
	"fmt"
	"gooze-vben-api/models"
	"mime/multipart"
	"slices"
	"strings"

	"github.com/soryetong/gooze-starter/gooze"
	"github.com/soryetong/gooze-starter/services/gzoss"
	"go.uber.org/zap"
)

func CheckIAVExt(ctx context.Context, types models.MaterialType, ext string) error {
	var allowExtList []string
	switch types {
	case models.MaterialType_Image:
		allowExtList = []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	case models.MaterialType_Audio:
		allowExtList = []string{".mp3", ".wav", ".aac"}
	case models.MaterialType_Video:
		allowExtList = []string{".mp4", ".mov", ".avi"}
	}
	if !slices.Contains(allowExtList, strings.ToLower(ext)) {
		return fmt.Errorf("不支持的文件类型")
	}

	return nil
}

// IAVSave 保存图片/音频/视频 types参数的目的是为了进行分层，不同场景使用不同的文件目录
func IAVSave(ctx context.Context, data *multipart.FileHeader, types string) (string, error) {
	ossType := gooze.Config.Oss.Type
	if ossType == "local" {
		types = "./static/" + types
	}
	result, err := gzoss.New(gzoss.OssType(ossType)).Upload(data, types)
	if err != nil {
		gooze.Log.WithCtx(ctx).Error("上传素材失败", zap.Error(err))
		return "", errors.New("上传素材失败")
	}

	return result.Url, nil
}
