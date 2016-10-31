package home

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/go-macaron/cache"
	"github.com/jicg/myblog/models"
	"gopkg.in/macaron.v1"
)

var (
	tpls = []string{"about", "index", "products", "life", "study"}
)

func Index(ctx *macaron.Context) {

	ctx.HTML(200, getTempletePath(ctx)+"index")
}

func Image(ctx *macaron.Context, cache cache.Cache) {
	//id = ctx.ParamsInt64(":id")
	filename := ctx.Params(":filename")
	uri := "/image/" + filename

	cachedata := cache.Get(uri)
	cachetype := cache.Get("img_" + filename + "_type")
	if cachedata != nil && cachetype != nil {
		datas, ok1 := cachedata.([]byte)
		filetype, ok2 := cachetype.(string)
		if ok1 && ok2 {
			ctx.Header().Add("content-type", "image/"+filetype)
			ctx.RawData(200, datas)
			return
		}
	}
	img := models.GetImageByURI(uri)
	if datas, err := ioutil.ReadFile(img.FilePath); err == nil {
		cache.Put("img_"+filename+"_type", img.FileType, 60*60*24*10)
		cache.Put(uri, datas, 60*60*24*10)
		ctx.Header().Add("content-type", "image/"+img.FileType)
		ctx.RawData(200, datas)
	} else {
		ctx.HTML(404, "404")
	}
}

func Any(ctx *macaron.Context) {
	name := ctx.ParamsEscape(":name")
	flag := false
	for i := 0; i < len(tpls); i++ {
		if strings.Compare(tpls[i], name) == 0 {
			flag = true
			break
		}
	}
	if name == "study" {
		paramstr := ctx.Params("*")
		params := strings.Split(paramstr, "/")
		cnt := 0
		ids := [2]int64{-1, -1}
		page := 1
		pagecnt := 6
		for i := len(params) - 1; i >= 0; i-- {
			str := params[i]
			if cnt < 2 {
				if id, e := strconv.ParseInt(str, 10, 64); e == nil {
					ids[cnt] = id
					cnt++
				}
			}
			if strings.Contains(str, "page") {
				s := strings.Replace(str, "page", "", -1)
				if p, e := strconv.ParseInt(s, 10, 0); e == nil {
					page = int(p)
				}
			}
			// if cnt == 2 {
			// 	break
			// }
		}
		ctx.Data["cnt"] = cnt
		ctx.Data["page"] = page
		cate := models.GetCategoryById(ids[0])
		if cnt == 1 && cate != nil && cate.Id != 0 {
			//得到当前类别下的文章
			count, e := models.GetArticleCountByCateId(ids[0])
			if e != nil {
				count = 0
			}
			arts := models.GetArticlePageByCateId(ids[0], page-1, pagecnt)
			ctx.Data["arts"] = arts
			ctx.Data["cateid"] = ids[0]
			ctx.Data["pagetotal"] = int64(math.Ceil(float64(count) / float64(pagecnt)))
		} else if cnt >= 2 {
			art := models.GetArticleById(ids[0])
			ctx.Data["art"] = art
			ctx.Data["cateid"] = ids[1]
		} else {
			count, e := models.GetAllArticleCount()
			if e != nil {
				count = 0
			}
			arts := models.GetAllArticlePage(page-1, pagecnt)
			ctx.Data["arts"] = arts
			ctx.Data["cateid"] = 0
			ctx.Data["pagetotal"] = int64(math.Ceil(float64(count) / float64(pagecnt)))
		}
		cates := models.GetAllCategoryOrderByOrderno()
		ctx.Data["cates"] = cates
	}
	if flag {
		ctx.HTML(200, getTempletePath(ctx)+name)
	} else {
		ctx.HTML(404, "404")
	}
}

func AddIdeal(ctx *macaron.Context) {
	title := ctx.QueryTrim("title")
	remark := ctx.QueryTrim("remark")
	if err := models.AddIdeal(title, remark); err != nil {
		ctx.Redirect("/error")
	}
	ctx.Redirect("/")
}

func getTempletePath(ctx *macaron.Context) string {
	return ctx.Data["templete_path"].(string)
}

func TempletePathHandel(ctx *macaron.Context) {
	ctx.Data["templete_path"] = "public/"
}
