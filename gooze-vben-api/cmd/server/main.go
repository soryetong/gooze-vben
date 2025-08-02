package main

import (
	_ "github.com/soryetong/gooze-starter/modules/casbinmodule"
	_ "github.com/soryetong/gooze-starter/modules/dbmodule"
	_ "gooze-vben-api/internal/bootstrap"

	"github.com/soryetong/gooze-starter/gooze"
)

func main() {
	gooze.Run()
}
