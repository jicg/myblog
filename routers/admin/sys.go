package admin

import (
	"gopkg.in/macaron.v1"
)

func GetSys(ctx *macaron.Context) {
	ctx.HTML(200, "admin/sys/index")
}
