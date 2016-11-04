package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jicg/myblog/modules/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

const XORM_LOG_PATH = "./data/log/xorm.log"

var x *xorm.Engine

type DBinfo struct {
	DBType string
	Host   string "127.0.0.1"
	Port   string "3306"
	User   string "root"
	Pwd    string "123"
	DBName string "myblog"
	Prefix string ""
}

func (d *DBinfo) loadConfig() {
	host := os.Getenv("BLOG_DB_HOST")
	if len(host) == 0 {
		sec := setting.Cfg.Section("database")
		d.Host = sec.Key("db_host").MustString("127.0.0.1")
		d.Port = sec.Key("db_port").MustString("3306")
		d.User = sec.Key("db_user").MustString("root")
		d.Pwd = sec.Key("db_pass").MustString("123")
		d.DBName = sec.Key("db_name").MustString("app_db")
		d.Prefix = sec.Key("dbprefix").MustString("")
	} else {
		d.Host = host

		port := os.Getenv("BLOG_DB_PORT")
		if len(port) != 0 {
			d.Port = port
		}

		user := os.Getenv("BLOG_DB_USER")
		if len(user) != 0 {
			d.User = user
		}
		d.User = user

		pwd := os.Getenv("BLOG_DB_PWD")
		if len(pwd) != 0 {
			d.Pwd = pwd
		}

		dbname := os.Getenv("BLOG_DB_DBNAME")
		if len(dbname) != 0 {
			d.DBName = dbname
		}

		prefix := os.Getenv("BLOG_DB_PREFIX")
		if len(prefix) != 0 {
			d.Prefix = os.Getenv("BLOG_DB_PREFIX")
		}
	}
}

func (d *DBinfo) DBUrl() string {
	fmt.Println(d.User + ":" + d.Pwd + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + "?charset=utf8")
	return d.User + ":" + d.Pwd + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + "?charset=utf8"
}

func init() {
	fmt.Println("＊＊＊＊＊＊＊＊＊数据库初始化＊＊＊＊＊＊＊＊＊＊＊")
	//得到配置信息
	dbinfo := DBinfo{DBType: "mysql", Host: "127.0.0.1", Port: "3306", User: "root", Pwd: "123", DBName: "myblog", Prefix: ""}
	dbinfo.loadConfig()
	//连接数据库
	var err error
	x, err = xorm.NewEngine(dbinfo.DBType, dbinfo.DBUrl())
	if err != nil {
		log.Println(err.Error(), dbinfo)
		return
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, dbinfo.Prefix)
	x.SetTableMapper(tbMapper)
	x.ShowSQL(true)
	x.Logger().SetLevel(core.LOG_INFO)
	os.MkdirAll("./data/log", 0777)
	f, err := os.Open(XORM_LOG_PATH)
	if err != nil {
		f, err = os.Create(XORM_LOG_PATH)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	//日志
	if err != nil {
		log.Fatalf("Fail to create xorm system logger: %v\n", err)
	}
	logger := xorm.NewSimpleLogger(f)
	logger.ShowSQL(true)
	x.SetLogger(logger)
	//缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	x.MapCacher(new(Nav), cacher)
	//engine.SetDefaultCacher(cacher)  engine.MapCacher(&user, nil)
	fmt.Println("＊＊＊＊＊＊＊＊＊同步表结构＊＊＊＊＊＊＊＊＊＊＊")

	err = x.Sync2(new(Nav), new(User), new(Ideal), new(Article), new(Category), new(Blog), new(Image))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("＊＊＊＊＊＊＊＊＊数据库加载成功＊＊＊＊＊＊＊＊＊＊＊")

}
