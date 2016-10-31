package admin

import (
	"io"
	"os"
	"strings"

	"github.com/jicg/myblog/models"
	"github.com/jicg/myblog/modules/util"
	"gopkg.in/macaron.v1"
)

var (
	filesize    int64
	filename    string
	newfilename string
	filepath    string
	filetype    string
	uri         string
)

type Size interface {
	Size() int64
}

func ListImage(ctx *macaron.Context) {
	imgs := models.GetAllImage()
	ctx.Data["imgs"] = imgs
	ctx.HTML(200, "admin/image/list")
}

func EditImage(ctx *macaron.Context) {
	var id = ctx.ParamsInt64(":id")
	img := models.GetImageById(id)
	ctx.Data["img"] = img
	ctx.HTML(200, "admin/image/edit")
}

func AddImage(ctx *macaron.Context) {
	name := ctx.QueryTrim("name")
	if err := InitFile(ctx); err != nil {
		ctx.Redirect("/admin/image/list")
		return
	}
	models.AddImage(name, uri, filepath, filetype, filename, filesize)
	ctx.Redirect("/admin/image/list")
}

func UpdateImage(ctx *macaron.Context) {
	var id = ctx.ParamsInt64(":id")
	img := models.GetImageById(id)
	name := ctx.QueryTrim("name")
	if err1 := InitFile(ctx); err1 != nil {
		panic(err1)
		ctx.Redirect("/admin/image/list")
		return
	}

	_, err2 := models.UpdateImage(id, name, img.Uri, filepath, filetype, filename, filesize)
	if err2 != nil {
		ctx.Redirect("/admin/image/list")
		return
	}

	if img.FilePath != filepath {
		os.Remove(img.FilePath)
	}
	ctx.Redirect("/admin/image/list")
}

func DelImage(ctx *macaron.Context) {
	id := ctx.ParamsInt64(":id")
	img := models.GetImageById(id)
	models.DeleteImageById(id)
	os.Remove(img.FilePath)
	ctx.Redirect("/admin/image/list")
}

func InitFile(ctx *macaron.Context) error {
	file, h, err1 := ctx.GetFile("imgfile")
	if err1 != nil {
		return err1
	}

	if sizeInterface, ok := file.(Size); ok {
		filesize = sizeInterface.Size()
	}
	filename = h.Filename
	//newfilename := hex.EncodeToString(com.RandomCreateBytes(16))
	newfilename = util.GetGuid()
	filepath = "./data/updateimage"
	if !util.IsDirExists(filepath) {
		os.Mkdir(filepath, 0777)
	}
	filepath = filepath + "/" + newfilename
	f, err2 := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		return err2
	}
	if _, err3 := io.Copy(f, file); err3 != nil {
		return err3
	}
	uri = "/image/" + newfilename
	index := strings.LastIndex(filename, ".") + 1
	filetype = "jpg"
	if filetype = filename[index:]; filetype == "jpeg" {
		filetype = "jpg"
	}
	return nil
}
