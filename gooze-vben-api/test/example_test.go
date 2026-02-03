package test

import (
	"context"
	"gooze-vben-api/internal/admin/dto"
	"gooze-vben-api/internal/admin/logic"
	"testing"

	"github.com/soryetong/gooze-starter/gooze"
	_ "github.com/soryetong/gooze-starter/modules/dbmodule"
)

func TestExample(t *testing.T) {
	gooze.RunTest("../configs/admin.yaml", "../.env.admin", false)

	// 注意：如果业务逻辑中涉及了 Redis、MySQL 等数据库操作，需要在 import 中匿名导入
	// 开始你的业务逻辑测试，如：
	params := &dto.DictListReq{
		Page:     1,
		PageSize: 10,
	}
	list, err := logic.NewSystemLogic().DictList(context.Background(), params)
	if err != nil {
		t.Fatal("请求异常: ", err)
	}

	t.Log("请求成功: ", list)
}
