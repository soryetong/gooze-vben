// This file was generated at: 2025-12-24 22:09:08
package bootstrap

import (
	"gooze-vben-api/internal/admin/logic"
	"gooze-vben-api/internal/admin/router"

	"github.com/soryetong/gooze-starter/gooze"
	"github.com/soryetong/gooze-starter/modules/httpmodule"
	"github.com/soryetong/gooze-starter/pkg/gzutil"
)

func init() {
	gooze.RegisterService(&AdminServer{})
}

type AdminServer struct {
	*gooze.IServer

	httpModule httpmodule.IHttp
}

func RegisterCustomHandlers() {
}

func (self *AdminServer) OnStart() (err error) {
	RegisterCustomHandlers()

	// 缓存api信息, 方便记录操作日志
	new(logic.SystemLogic).CacheApiInfo()

	// 添加回调函数
	self.httpModule.OnStop(self.exitCallback())

	self.httpModule.Init(self, gooze.Config.App.Addr, gooze.Config.App.Timeout, router.InitRouter())
	err = self.httpModule.Start()

	return
}

// TODO 添加回调函数, 无逻辑可直接删除这个方法
func (self *AdminServer) exitCallback() *gzutil.OrderlyMap {
	callback := gzutil.NewOrderlyMap()
	callback.Append("exit", func() {
		gooze.Log.Info("这是程序退出后的回调函数, 执行你想要执行的逻辑, 无逻辑可以直接删除这段代码")
	})

	return callback
}
