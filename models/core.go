package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jicg/myblog/modules/setting"
)

const XORM_LOG_PATH = "./data//log/xorm.log"

var x *xorm.Engine

func init() {
	fmt.Println("＊＊＊＊＊＊＊＊＊数据库初始化＊＊＊＊＊＊＊＊＊＊＊")
	//得到配置信息
	sec := setting.Cfg.Section("database")
	dbhost := sec.Key("db_host").MustString("127.0.0.1")
	dbport := sec.Key("db_port").MustString("3306")
	dbuser := sec.Key("db_user").MustString("root")
	dbpassword := sec.Key("db_pass").MustString("123")
	dbname := sec.Key("db_name").MustString("app_db")
	dbprefix := sec.Key("dbprefix").MustString("")
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	//连接数据库
	var err error
	x, err = xorm.NewEngine("mysql", dburl)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, dbprefix)
	x.SetTableMapper(tbMapper)
	x.ShowSQL(true)
	x.Logger().SetLevel(core.LOG_INFO)
	//x.Logger().SetLevel(core.LOG_DEBUG)
	//logWriter, err := syslog.New(syslog.LOG_DEBUG, "xorm")
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
