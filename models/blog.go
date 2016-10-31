package models

import (
	"time"
)

type Blog struct {
	Id      int64 `xorm:"pk autoincr"`
	Pid     int64
	Title   string
	Remark  string    `xorm:"text"`
	Created time.Time `xorm:"created"`
}

func (blog *Blog) Save() error {
	_, err := x.Insert(blog)
	return err
}

func AddBlog(pid int64, title string, remark string) error {
	_, err := x.Insert(&Blog{Title: title, Pid: pid, Remark: remark})
	return err
}

func (blog *Blog) Update() error {
	_, err := x.Update(blog)
	return err
}

func GetAllBlog() *[]Blog {
	blogs := new([]Blog)
	x.Find(blogs)
	return blogs
}
