package models

type MaterialType string

const (
	MaterialType_Image MaterialType = "image"
	MaterialType_Audio MaterialType = "audio"
	MaterialType_Video MaterialType = "video"
)

// CMaterials 素材管理
type CMaterials struct {
	GnModel
	Name        string       `json:"name" gorm:"name"`                // 素材名称
	ResourceUrl string       `json:"resourceUrl" gorm:"resource_url"` // 素材地址
	Description string       `json:"description" gorm:"description"`  // 描述
	Status      int64        `json:"status" gorm:"status"`            // 状态（1-显示 2-隐藏）
	Type        MaterialType `json:"type" gorm:"type"`                // 素材类型
}

// TableName 表名称
func (*CMaterials) TableName() string {
	return "c_materials"
}
