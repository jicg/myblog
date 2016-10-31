package admin

import (
	"github.com/go-macaron/session"
	"github.com/jicg/myblog/models"
	"gopkg.in/macaron.v1"
)

type NavTreeChildren struct {
	Id      int64  `json:"id"`
	Pid     int64  `json:"pid"`
	Orderno int64  `json:"orderno"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Icon    string `json:"icon"`
	Issys   bool   `json:"issys"`
}

type NavTree struct {
	Id      int64              `json:"id"`
	Pid     int64              `json:"pid"`
	Orderno int64              `json:"orderno"`
	Name    string             `json:"name"`
	Url     string             `json:"url"`
	Icon    string             `json:"icon"`
	Issys   bool               `json:"issys"`
	Navs    []*NavTreeChildren `json:"navs"`
}

func ListNav(ctx *macaron.Context) {
	//ctx.Data["navs"] = models.GetPageNavByPid(0, 10, 0)
	tempChildrens := make(map[int64][]*NavTreeChildren)
	navsTree := make([]*NavTree, 0)
	navs := models.GetAllNavOrderByOrderno()
	for _, v := range navs {
		if v.Pid == 0 {
			navsTree = append(navsTree, &NavTree{
				Id:      v.Id,
				Pid:     v.Pid,
				Orderno: v.Orderno,
				Name:    v.Name,
				Issys:   v.Issys,
				Url:     v.Url, Icon: v.Icon})
		} else {
			tempChildren := tempChildrens[v.Pid]
			if tempChildren == nil {
				tempChildren = make([]*NavTreeChildren, 0)
			}
			tempChildrens[v.Pid] = append(tempChildren, &NavTreeChildren{
				Id:      v.Id,
				Pid:     v.Pid,
				Orderno: v.Orderno,
				Name:    v.Name,
				Url:     v.Url,
				Issys:   v.Issys,
				Icon:    v.Icon})
		}
	}
	for _, v := range navsTree {
		v.Navs = tempChildrens[v.Id]
	}
	ctx.Data["navs"] = navsTree
	/*if bytes, err := json.Marshal(navs); err == nil {
		fmt.Println(string(bytes))
	}*/
	ctx.HTML(200, "admin/sys/navs/list")
}
func AddNav(ctx *macaron.Context, sess session.Store) {
	orderno := ctx.QueryInt64("orderno")
	name := ctx.Query("name")
	url := ctx.Query("url")
	pid := ctx.QueryInt64("pid")
	models.SaveNav(name, url, orderno, pid) //name string, url string, orderno int64, pid int64
	ctx.Redirect("/admin/sys/navs/list")
}

func UpdateNav(ctx *macaron.Context, sess session.Store) {
	id := ctx.QueryInt64("id")
	orderno := ctx.QueryInt64("orderno")
	name := ctx.Query("name")
	url := ctx.Query("url")
	pid := ctx.QueryInt64("pid")
	models.UpdateNav(id, name, url, orderno, pid)
	ctx.Redirect("/admin/sys/navs/list")
}

func EditNav(ctx *macaron.Context) {
	var id = ctx.ParamsInt64(":id")
	nav := models.GetNavById(id)
	ctx.Data["nav"] = nav
	if nav.Id != 0 {
		if cnt, err := models.GetNavCountByPid(nav.Id); err == nil {
			ctx.Data["chcnt"] = cnt
		} else {
			ctx.Data["chcnt"] = 0
		}
	}
	ctx.Data["topnavs"] = models.GetPageNavByPid(0, 100, 0)
	ctx.HTML(200, "admin/sys/navs/edit")
}

func DelNav(ctx *macaron.Context) {
	id := ctx.ParamsInt64(":id")
	nav := models.GetNavById(id)
	if nav.Pid == 0 {
		models.DeleteNavByPid(id)
		models.DeleteNavById(id)
	} else {
		models.DeleteNavById(id)
	}
	ctx.Redirect("/admin/sys/navs/list")
}
