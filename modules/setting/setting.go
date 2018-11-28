package setting

import (
	"fmt"
	"html/template"
	"strings"

	"os"

	"github.com/Unknwon/log"
	"github.com/go-macaron/cache"
	_ "github.com/go-macaron/cache/ledis"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/session"
	_ "github.com/go-macaron/session/ledis"
	"gopkg.in/ini.v1"
	"gopkg.in/macaron.v1"
)

const (
	CACHE_DIR   = "./data/cache.db,db=0"
	SESSION_DIR = "./data/session.db,db=0"
)

var (
	ProdMode bool
	// Server settings.
	HTTPPort int
	// Global settings.
	Cfg *ini.File

	IsInstall bool
)

func init() {
	HTTPPort = 8080
	if _, err := os.Open("conf/app.ini"); err != nil {
		if _, e := os.Create("conf/app.ini"); e != nil {
			log.FatalD(4, "Fail to create file: %v", e)
		}

	}
	sources := []interface{}{"sys.ini", "conf/app.ini"}
	var err error
	Cfg, err = macaron.SetConfig(sources[0], sources[1:]...)
	if err != nil {
		log.FatalD(4, "Fail to set configuration: %v", err)
	}
	sec := Cfg.Section("server")
	HTTPPort = sec.Key("httpport").MustInt(8080)
	log.Prefix = sec.Key("appname").MustString("app")
	ProdMode = sec.Key("runmode").MustString("app") != "dev"
}

func NewCache() macaron.Handler {
	var op cache.Options
	op = cache.Options{
		Adapter:       "ledis",
		AdapterConfig: "data_dir=" + CACHE_DIR,
	}
	return cache.Cacher(op)
}

func NewStatic(dir string, pre string) macaron.Handler {
	return macaron.Static(dir, macaron.StaticOptions{
		// 请求静态资源时的 URL 前缀，默认没有前缀
		Prefix: dir,
		// 禁止记录静态资源路由日志，默认为不禁止记录
		SkipLogging: true,
		// 当请求目录时的默认索引文件，默认为 "index.html"
		//IndexFile: "index.html",
		// 用于返回自定义过期响应头，默认为不设置
		// https://developers.google.com/speed/docs/insights/LeverageBrowserCaching
		//Expires: func() string {
		//	return time.Now().Add(24 * 60 * time.Minute).UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
		//},
	})
}

func NewCap() macaron.Handler {
	return captcha.Captchaer(captcha.Options{
		// 获取验证码图片的 URL 前缀，默认为 "/captcha/"
		URLPrefix: "/captcha/",
		// 表单隐藏元素的 ID 名称，默认为 "captcha_id"
		FieldIdName: "captcha_id",
		// 用户输入验证码值的元素 ID，默认为 "captcha"
		FieldCaptchaName: "captcha",
		// 验证字符的个数，默认为 6
		ChallengeNums: 6,
		// 验证码图片的宽度，默认为 240 像素
		Width: 240,
		// 验证码图片的高度，默认为 80 像素
		Height: 80,
		// 验证码过期时间，默认为 600 秒
		Expiration: 600,
		// 用于存储验证码正确值的 Cache 键名，默认为 "captcha_"
		CachePrefix: "hw_",
	})
}

func NewSession() macaron.Handler {
	return session.Sessioner(session.Options{
		// 提供器的名称，默认为 "memory"
		Provider: "ledis",
		// 提供器的配置，根据提供器而不同
		ProviderConfig: "data_dir=" + SESSION_DIR,
		// 用于存放会话 ID 的 Cookie 名称，默认为 "MacaronSession"
		CookieName: "AppSession",
		// Cookie 储存路径，默认为 "/"
		CookiePath: "/",
		// GC 执行时间间隔，默认为 3600 秒
		Gclifetime: 3600,
		// 最大生存时间，默认和 GC 执行时间间隔相同
		Maxlifetime: 3600,
		// 仅限使用 HTTPS，默认为 false
		Secure: false,
		// Cookie 生存时间，默认为 0 秒
		CookieLifeTime: 0,
		// Cookie 储存域名，默认为空
		Domain: "",
		// 会话 ID 长度，默认为 16 位
		IDLength: 16,
		// 配置分区名称，默认为 "session"
		Section: "session",
	})
}

func NewRender() macaron.Handler {
	return macaron.Renderer(macaron.RenderOptions{
		// 模板文件目录，默认为 "templates"
		Directory: "views",
		// Specify a layout template. Layouts can call {{ yield }} to render the current template.
		//Layout: "layout",
		// 模板文件后缀，默认为 [".tmpl", ".html"]
		Extensions: []string{".tmpl", ".html"},
		// 模板函数，默认为 []
		Funcs: []template.FuncMap{map[string]interface{}{
			"AppName": func() string {
				return "Macaron"
			},
			"AppVer": func() string {
				return "1.0.0"
			},
			"equals":    equals,
			"notequals": notequals,
			"big":       big,
			"empty":     empty,
			"notempty":  notempty,
		}},
		// 模板语法分隔符，默认为 ["{{", "}}"]
		Delims: macaron.Delims{"{{", "}}"},
		// 追加的 Content-Type 头信息，默认为 "UTF-8"
		Charset: "UTF-8",
		// 渲染具有缩进格式的 JSON，默认为不缩进
		IndentJSON: false,
		// 渲染具有缩进格式的 XML，默认为不缩进
		IndentXML: false,
		// 渲染具有前缀的 JSON，默认为无前缀
		//PrefixJSON: []byte("jicg_app"),
		// 渲染具有前缀的 XML，默认为无前缀
		//PrefixXML: []byte("jicg_app"),
		// 允许输出格式为 XHTML 而不是 HTML，默认为 "text/html"
		HTMLContentType: "text/html",
	})
}

func equals(a, b interface{}) (equal bool) {
	equal = false
	if strings.TrimSpace(fmt.Sprintf("%v", a)) == strings.TrimSpace(fmt.Sprintf("%v", b)) {
		equal = true
	}
	return
}

func notequals(a, b interface{}) (equal bool) {
	equal = true
	if strings.TrimSpace(fmt.Sprintf("%v", a)) == strings.TrimSpace(fmt.Sprintf("%v", b)) {
		equal = false
	}
	return
}

func big(a int64, b int64) (flag bool) {
	flag = a > b
	return
}

func empty(a interface{}) (flag bool) {
	flag = a == nil
	return
}

func notempty(a interface{}) (flag bool) {
	flag = a != nil
	return
}
