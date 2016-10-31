package admin

import (
	"encoding/json"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/session"
	"github.com/jicg/myblog/models"
	"gopkg.in/macaron.v1"
)

func LoadNavBars(ctx *macaron.Context, cache cache.Cache, sess session.Store) {

	url := models.FmtUrl(ctx.Req.URL.Path)
	var navs []*models.Nav
	var navsobj = cache.Get("navs")
	if navsobj == nil {
		navs = models.GetAllNavOrderByOrderno()
		if bytes, err := json.Marshal(navs); err == nil {
			cache.Put("navs", string(bytes), 0)
		}
	} else {
		json.Unmarshal([]byte(navsobj.(string)), &navs)
	}
	if navs == nil || len(navs) == 0 {
		return
	}
	if nav, err := getNavByUrl(url, navs); err == nil && nav != nil {
		mean := make(map[string]interface{})
		mean["navbars"], mean["sidebars"], mean["id"], mean["pid"] = getMeanByidAndPid(nav.Id, nav.Pid, navs)
		sess.Set("MEAN", mean)
		ctx.Data["MEAN"] = mean
	} else {
		if sess.Get("MEAN") == nil {
			mean := make(map[string]interface{})
			mean["navbars"], mean["sidebars"], mean["id"], mean["pid"] = getMeanByidAndPid(navs[0].Id, navs[0].Pid, navs)
			ctx.Data["MEAN"] = mean
		} else {
			ctx.Data["MEAN"] = sess.Get("MEAN").(map[string]interface{})
		}

	}

}

func getNavByUrl(url string, navs []*models.Nav) (*models.Nav, error) {
	for _, val := range navs {
		if url == val.Url {
			return val, nil
		}
	}
	return nil, models.ErrNavNotExist{Errstr: "会员不存在"}
}

func getMeanByidAndPid(id int64, pid int64, navs []*models.Nav) ([]*models.Nav, []*models.Nav, int64, int64) {
	if pid <= 0 {
		pid = id
	}
	navbars := make([]*models.Nav, 0)
	sidebars := make([]*models.Nav, 0)
	for _, val := range navs {
		if val.Pid == pid {
			sidebars = append(sidebars, val)
		}
		if val.Pid == 0 {
			navbars = append(navbars, val)
		}
	}
	return navbars, sidebars, id, pid
}
