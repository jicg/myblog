package main

import (
	"encoding/gob"

	_ "github.com/jicg/myblog/modules/setting"

	"github.com/jicg/myblog/modules/setting"
	"github.com/jicg/myblog/routers/admin"
	"github.com/jicg/myblog/routers/comm"
	"github.com/jicg/myblog/routers/home"

	"github.com/jicg/myblog/models"

	"net/http"

	"strconv"

	"gopkg.in/macaron.v1"
)

var (
	m *macaron.Macaron
)

//main入口
func main() {
	newMacaron()
	m.Group("", indexGroup, home.TempletePathHandel)
	//m.Run(setting.HTTPPort)

	http.Handle("/", m)
	http.Handle("/image/", &comm.ImageHandle{})
	http.ListenAndServe(":"+strconv.Itoa(setting.HTTPPort), nil)
}

//初始化macaron
func newMacaron() {
	gob.Register([]*models.Nav{})
	gob.Register(&models.User{})
	if m == nil {
		//macaron.Env = macaron.PROD
		m = macaron.New()
		m.Use(macaron.Logger())
		m.Use(macaron.Recovery())
		m.Use(setting.NewStatic("static", "static"))
		m.Use(setting.NewStatic("data/update", "update"))
		m.Use(setting.NewCache())
		m.Use(setting.NewCap())
		m.Use(setting.NewRender())
		m.Use(setting.NewSession())
		macaron.Config().SectionStrings()
	}
}

//首页
func indexGroup() {
	m.Get("/", home.Index)
	m.Get("/image/:filename", home.Image)
	m.Any("/user/login", admin.Login)
	m.Get("/user/logout", admin.LoginCheck, admin.Logout)
	m.Get("/user/clrSess", admin.LoginCheck, admin.ClearSessionCache)
	m.Get("/user/clrCache", admin.LoginCheck, admin.ClearCache)

	m.Group("/admin", adminGroup, admin.LoginCheck, admin.LoadNavBars)
	m.Group("/ideal", func() {
		m.Post("/add", home.AddIdeal)
	})
	m.Get("/:name/?*", home.Any)
}

//后台*****************************************
func adminGroup() {
	m.Group("", adminIndexGroup)
	m.Group("/index", adminIndexGroup)
	m.Group("/article", adminActicleGroup)
	m.Group("/category", adminCategoryGroup)
	m.Group("/image", adminImageGroup)
	m.Group("/sys", adminSysGroup)
	m.Group("/*", adminAnyGroup)
}

//后台->首页
func adminIndexGroup() {
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "admin/index")
	})
}

//后台->系统设置
func adminSysGroup() {
	m.Get("/", admin.GetSys)
	m.Get("/index", admin.GetSys)

	m.Group("/param", func() {
		m.Get("/list", admin.ListParam)
		m.Get("/edit/:id", admin.EditParam)
		m.Post("/add", admin.AddParam)
		m.Post("/update/:id", admin.UpdateParam)
		m.Any("/del/:id", admin.DelParam)
	})

	m.Group("/navs", func() {
		m.Get("/list", admin.ListNav)
		m.Get("/edit/:id", admin.EditNav)
		m.Post("/add", admin.AddNav)
		m.Post("/update/:id", admin.UpdateNav)
		m.Any("/del/:id", admin.DelNav)
	})
}

func adminActicleGroup() {
	m.Get("/list", admin.ListActicle)
	m.Get("/edit/:id/?:edittype", admin.EditActicle)
	m.Post("/add", admin.AddActicle)
	m.Post("/update/:id", admin.UpdateActicle)
	m.Get("/del/:id", admin.DelActicle)
}
func adminCategoryGroup() {
	m.Get("/list", admin.ListCategory)
	m.Get("/edit/:id", admin.EditCategory)
	m.Post("/add", admin.AddCategory)
	m.Post("/update/:id", admin.UpdateCategory)
	m.Get("/del/:id", admin.DelCategory)
}

func adminImageGroup() {
	m.Get("/list", admin.ListImage)
	m.Get("/edit/:id", admin.EditImage)
	m.Post("/add", admin.AddImage)
	m.Post("/update/:id", admin.UpdateImage)
	m.Get("/del/:id", admin.DelImage)
}

//后台－>未定义界面
func adminAnyGroup() {
	m.Get("", func(ctx *macaron.Context) {
		ctx.Data["html"] = "页面构思中。。。。。"
		ctx.HTML(200, "admin/any")
	})
}
