package admin

import (
	"github.com/jicg/myblog/models"
	"gopkg.in/macaron.v1"
)

func ListCategory(ctx *macaron.Context) {
	cates := models.GetAllCategoryOrderByOrderno()
	ctx.Data["cates"] = cates
	ctx.HTML(200, "admin/category/list")
}

func EditCategory(ctx *macaron.Context) {
	var id = ctx.ParamsInt64(":id")
	cate := models.GetCategoryById(id)
	ctx.Data["cate"] = cate
	ctx.HTML(200, "admin/category/edit")
}

func AddCategory(ctx *macaron.Context) {
	orderno := ctx.QueryInt64("orderno")
	name := ctx.Query("name")
	desc := ctx.Query("desc")
	models.SaveCategory(name, desc, orderno)
	ctx.Redirect("/admin/category/list")
}

func UpdateCategory(ctx *macaron.Context) {
	id := ctx.QueryInt64("id")
	orderno := ctx.QueryInt64("orderno")
	name := ctx.Query("name")
	desc := ctx.Query("desc")
	// pid := ctx.QueryInt64("pid")
	// artid := ctx.QueryInt64("artid")
	// models.UpdateCategory(id, name, desc, orderno, pid, artid)
	models.UpdateCategory(id, name, desc, orderno)
	ctx.Redirect("/admin/category/list")
}

func DelCategory(ctx *macaron.Context) {
	id := ctx.ParamsInt64(":id")
	models.DeleteCategoryById(id)
	ctx.Redirect("/admin/category/list")
}
