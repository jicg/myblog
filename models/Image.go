package models

import (
	"time"
)

type Image struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"    xorm:"unique index"`
	Uri      string    `json:"uri"     xorm:"unique index"`
	FilePath string    `json:"filepath"`
	FileType string    `json:"filetype"`
	FileName string    `json:"filename"`
	FileSize int64     `json:"filesize"`
	Created  time.Time `json:"create"    xorm:"created"`
	Updated  time.Time `json:"updated"   xorm:"updated"`
}

func GetAllImage() []*Image {
	imgs := make([]*Image, 0)
	x.Asc("id").Find(&imgs)
	return imgs
}

func GetImageById(id int64) *Image {
	img := new(Image)
	x.Where("id = ? ", id).Get(img)
	return img
}

func AddImage(name string, uri string, filepath string, filetype string, filename string, filesize int64) (int64, error) {
	return x.InsertOne(&Image{Name: name, Uri: uri, FilePath: filepath, FileType: filetype, FileName: filename, FileSize: filesize})
}

func UpdateImage(id int64, name string, uri string, filepath string, filetype string, filename string, filesize int64) (int64, error) {
	return x.Where(" id = ? ", id).Update(&Image{Name: name, Uri: uri, FilePath: filepath, FileType: filetype, FileName: filename, FileSize: filesize})
}

func GetImageByURI(uri string) *Image {
	img := new(Image)
	x.Where(" uri = ? ", uri).Get(img)
	return img
}

func DeleteImageById(id int64) (int64, error) {
	return x.Where("id = ?", id).Delete(new(Image))
}
