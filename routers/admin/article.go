package admin

import (
	"strconv"

	"github.com/jicg/myblog/models"
	"gopkg.in/macaron.v1"
)

var (
	edittype int8
	content  string
	source   string
)

func ListActicle(ctx *macaron.Context) {
	arts := models.GetAllArticleOrerByID()
	ctx.Data["arts"] = arts
	ctx.HTML(200, "admin/article/list")
}

func EditActicle(ctx *macaron.Context) {
	var id = ctx.ParamsInt64(":id")
	pedittype := ctx.ParamsInt64(":edittype")
	art := models.GetArticleById(id)
	ctx.Data["art"] = art
	edittype1 := int8(1)
	if pedittype > 0 {
		edittype1 = int8(pedittype)
	} else {
		if art.Id == 0 {
			user := ctx.Data["user"].(*models.User)
			edittype1 = user.EditorType
		} else {
			edittype1 = art.ArtType
		}
	}
	cates := models.GetAllCategoryOrderByOrderno()
	ctx.Data["cates"] = cates
	ctx.Data["edittype"] = edittype1
	ctx.HTML(200, "admin/article/edit")
}

func AddActicle(ctx *macaron.Context) {
	name := ctx.Query("name")
	desc := ctx.Query("desc")
	cateid := ctx.QueryInt64("cateid")
	selectEditWay(ctx)

	models.SaveActicle(name, cateid, desc, content, edittype, source)
	ctx.Redirect("/admin/article/list")
}

func UpdateActicle(ctx *macaron.Context) {
	var id = ctx.ParamsInt64(":id")
	name := ctx.Query("name")
	desc := ctx.Query("desc")
	cateid := ctx.QueryInt64("cateid")
	selectEditWay(ctx)
	models.UpdateActicle(id, cateid, name, desc, content, edittype, source)
	ctx.Redirect("/admin/article/list")
}

func selectEditWay(ctx *macaron.Context) error {
	if v, err := strconv.ParseUint(ctx.QueryTrim("edittype"), 10, 8); err != nil {
		edittype = int8(1)
	} else {
		edittype = int8(v)
	}
	if edittype == 1 {
		source = ctx.Query("content")
		content = ctx.Query("content")
	} else {
		source = ctx.Query("content")
		content = ctx.Query("content-div-html-code")
	}
	return nil
}

func DelActicle(ctx *macaron.Context) {
	id := ctx.ParamsInt64(":id")
	models.DeleteActicleById(id)
	ctx.Redirect("/admin/article/list")
}
