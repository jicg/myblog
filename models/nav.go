package models

import (
	"strings"
	"time"
)

type Nav struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"      xorm:"index"`
	Url     string    `json:"url"       xorm:"unique index"`
	Icon    string    `json:"icon"`
	Orderno int64     `json:"orderno"`
	Pid     int64     `json:"pid"`
	Issys   bool      `json:"issys"     xorm:"Bool"`
	Status  bool      `json:"status"    xorm:"Bool"`
	Created time.Time `json:"create"    xorm:"created"`
	Updated time.Time `json:"updated"   xorm:"updated"`
	Delete  time.Time `json:"delete"    xorm:"deleted"`
}

func UpdateNav(id int64, name string, url string, orderno int64, pid int64) (int64, error) {
	return x.Id(id).Update(&Nav{Name: name, Url: url, Orderno: orderno, Pid: pid, Issys: false, Status: true})
}

func SaveNav(name string, url string, orderno int64, pid int64) (int64, error) {
	return x.InsertOne(&Nav{Name: name, Url: url, Orderno: orderno, Pid: pid, Issys: false, Status: true})
}

func DeleteNavById(id int64) (int64, error) {
	return x.Where("id = ?", id).Delete(new(Nav))
}

func DeleteNavByPid(pid int64) (int64, error) {
	return x.Where("pid = ?", pid).Delete(new(Nav))
}

func GetNavById(id int64) *Nav {
	nav := new(Nav)
	x.Id(id).Get(nav)
	return nav
}

//静态方法
func GetAllNavOrderByPid() []*Nav {
	navs := make([]*Nav, 0)
	x.Where("status = ?", true).Asc("pid").Find(&navs)
	return navs
}

func GetAllNavOrderByOrderno() []*Nav {
	navs := make([]*Nav, 0)
	x.Where("status = ?", true).Asc("orderno").Find(&navs)
	return navs
}

func GetNavCountByPid(id int64) (int64, error) {
	return x.Where("pid = ?", id).Where("status = ?", true).Count(new(Nav))
}

func GetAllNavCount() (int64, error) {
	return x.Where("status = ?", true).Count(new(Nav))
}

//sidebar.Query().Filter("isactive", "Y").Filter("pid", pid).Limit(pagecnt, (page-1)*pagecnt).All(&sidebars)
func GetPageNavByPid(pid int64, pagecnt int, start int) []*Nav {
	navs := make([]*Nav, 0)
	x.Where("status=?", true).And("pid=?", pid).Limit(pagecnt, start).Find(&navs)
	return navs
}

func GetNavByUrl(url string) (*Nav, error) {
	nav := &Nav{Url: url}
	has, err := x.Get(nav)
	if err != nil {
		return nil, err
	}
	if has {
		return nav, nil
	}
	return nil, ErrNavNotExist{Errstr: "会员不存在"}
}

func FmtUrl(url string) string {
	return "/" + strings.Trim(url, "/")
}
