package admin

import (
	"gopkg.in/macaron.v1"
)

func ListParam(ctx *macaron.Context) {
	ctx.HTML(200, "admin/sys/param/list")
}

func EditParam(ctx *macaron.Context) {
	ctx.HTML(200, "admin/sys/param/list")
}
func AddParam(ctx *macaron.Context) {
	ctx.HTML(200, "admin/sys/param/list")
}
func UpdateParam(ctx *macaron.Context) {
	ctx.HTML(200, "admin/sys/param/list")
}

func DelParam(ctx *macaron.Context) {
	ctx.HTML(200, "admin/sys/param/list")
}
