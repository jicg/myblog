package models

import (
	"time"
)

type Ideal struct {
	Id      int64 `xorm:"pk autoincr"`
	Title   string
	Remark  string    `xorm:"text"`
	Created time.Time `xorm:"created"`
}

func (ideal *Ideal) Save() error {
	_, err := x.Insert(ideal)
	return err
}

func AddIdeal(title string, remark string) error {
	_, err := x.Insert(&Ideal{Title: title, Remark: remark})
	return err
}

func GetAllIdeal() *[]Ideal {
	ideals := new([]Ideal)
	x.Find(ideals)
	return ideals
}
