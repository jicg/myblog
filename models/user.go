package models

import (
	"time"
)

type User struct {
	Id         int64
	Name       string `xorm:"index unique"`
	Pwd        string
	Email      string    `xorm:"index unique"`
	Nickname   string    `xorm:"index"`
	EditorType int8      //1 cke ; 2 editormd
	Issys      bool      `xorm:"Bool"`
	Status     bool      `xorm:"Bool"`
	Created    time.Time `xorm:"created"`
	Updated    time.Time `xorm:"updated"`
}

func GetUserByNameWithPwd(name string, pwd string) *User {
	user := new(User)
	x.Where(" name = ?", name).And(" pwd = ? ", pwd).Get(user)
	return user
}
