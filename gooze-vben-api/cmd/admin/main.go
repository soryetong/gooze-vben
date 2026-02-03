package main

import (
	_ "gooze-vben-api/internal/admin/bootstrap"

	_ "github.com/soryetong/gooze-starter/modules/casbinmodule"
	_ "github.com/soryetong/gooze-starter/modules/dbmodule"

	"github.com/soryetong/gooze-starter/gooze"
)

func main() {
	gooze.Run()
}
