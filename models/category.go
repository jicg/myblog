package models

import (
	"time"
)

type Category struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"      xorm:"unique index"`
	Desc    string    `json:"desc"`
	Orderno int64     `json:"orderno"`
	Created time.Time `json:"create"    xorm:"created"`
	Updated time.Time `json:"updated"   xorm:"updated"`
}

func (cate *Category) Save() error {
	_, err := x.Insert(cate)
	return err
}

func SaveCategory(name string, desc string, orderno int64) (int64, error) {
	return x.InsertOne(&Category{Name: name, Desc: desc, Orderno: orderno})
}

func UpdateCategory(id int64, name string, desc string, orderno int64) (int64, error) {
	return x.Id(id).Update(&Category{Name: name, Desc: desc, Orderno: orderno})
}

func GetAllCategoryOrderByOrderno() []*Category {
	cates := make([]*Category, 0)
	x.Asc("orderno").Find(&cates)
	return cates
}

func GetCategoryById(id int64) *Category {
	cate := new(Category)
	x.Where("id = ? ", id).Get(cate)
	return cate
}

func DeleteCategoryById(id int64) (int64, error) {
	return x.Where("id = ?", id).Delete(new(Category))
}
