package admin

import (
	"github.com/go-macaron/session"
	"github.com/jicg/myblog/models"
	"gopkg.in/macaron.v1"
)

func Login(ctx *macaron.Context, sess session.Store) {
	user := sess.Get("USER")
	if user == nil {
		name := ctx.QueryTrim("username")
		pwd := ctx.QueryTrim("password")
		if len(name) != 0 && len(pwd) != 0 {
			user := models.GetUserByNameWithPwd(name, pwd)
			if user.Id == 0 {
				ctx.Data["error"] = "用户名或密码错误"
			} else {
				sess.Set("USER", user)
				ctx.Redirect("/admin")
				return
			}
		}
	} else {
		ctx.Data["user"] = user
	}
	ctx.HTML(200, "admin/login")
}

func Logout(ctx *macaron.Context, sess session.Store) {
	sess.Destory(ctx)
	ctx.Redirect("/user/login")
}

func LoginCheck(ctx *macaron.Context, sess session.Store) {
	user := sess.Get("USER")
	//u, ok := user.(models.User) || ok && u.Id == 0
	if user == nil {
		ctx.Redirect("/user/login")
	} else {
		ctx.Data["user"] = user
	}
}
