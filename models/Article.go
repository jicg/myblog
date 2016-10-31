package models

import (
	"html/template"
	"time"
)

type Article struct {
	Id       int64         `json:"id"        xorm:"pk autoincr"`
	Name     string        `json:"name"      xorm:"unique index"`
	Desc     string        `json:"desc"`
	CateId   int64         `json:cateid     xorm:"index"`
	CateName string        `json:catename`
	ArtType  int8          ////arttype 1 html ，2 makedown
	Content  template.HTML `json:"content"   xorm:"text"`
	Source   string        `'json:"source"  xorm:"text"`
	Created  time.Time     `json:"create"    xorm:"created"`
	Updated  time.Time     `json:"updated"   xorm:"updated"`
}

func GetAllArticleOrerByID() *[]Article {
	arts := new([]Article)
	x.Find(arts)
	return arts
}

//GetAllArticlePage 得到当前页面的文章
func GetAllArticlePage(page int, pagecnt int) *[]Article {
	arts := new([]Article)
	x.Desc("updated").Limit(pagecnt, page).Find(arts)
	return arts
}

func GetArticleById(id int64) *Article {
	art := new(Article)
	x.Id(id).Get(art)
	return art
}

//GetArticleByCateId 得到该类别下当前页面的文章
func GetArticleByCateId(cateid int64) *[]Article {
	arts := new([]Article)
	x.Where("cate_id = ?", cateid).Find(arts)
	return arts
}

//GetArticlePageByCateId 得到该类别下当前页面的文章
func GetArticlePageByCateId(cateid int64, page int, pagecnt int) *[]Article {
	arts := new([]Article)
	x.Where("cate_id = ?", cateid).Limit(pagecnt, page).Find(arts)
	return arts
}

//GetArticleCountByCateId 得到该类别下的文章总数量
func GetArticleCountByCateId(cateid int64) (int64, error) {
	return x.Where("cate_id = ?", cateid).Count(new(Article))
}

//GetAllArticleCount 得到总数量
func GetAllArticleCount() (int64, error) {
	return x.Count(new(Article))
}

func SaveActicle(name string, cateid int64, desc string, content string, arttype int8, source string) (int64, error) {
	cate := new(Category)
	x.Id(cateid).Get(cate)
	return x.InsertOne(&Article{Name: name, CateId: cateid, CateName: cate.Name, Desc: desc, Content: template.HTML(content), ArtType: arttype, Source: source})
}

func UpdateActicle(id int64, cateid int64, name string, desc string, content string, arttype int8, source string) (int64, error) {
	cate := new(Category)
	x.Id(cateid).Get(cate)
	return x.Id(id).Update(&Article{Name: name, CateId: cateid, CateName: cate.Name, Desc: desc, Content: template.HTML(content), ArtType: arttype, Source: source})
}

func DeleteActicleById(id int64) (int64, error) {
	return x.Where("id = ?", id).Delete(new(Article))
}
