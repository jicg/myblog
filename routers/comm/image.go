package comm

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/twinj/uuid"

	"path/filepath"
	"strings"

	"errors"
)

const (
	// IMG_DIR 默认路径
	IMG_DIR string = "./data/files"
)

type ImageHandle struct {
	imgdir  string
	MaxSize int64
}

// ImageHandleInfo 文本上传类型定义
type ImageHandleInfo struct {
	Flag     string
	Fromname string
	ToData   func(data *RetData) []byte
}

//RetData 定义vo 运行结果存储该对象，在转换[]byte
type RetData struct {
	Code     int64
	Msg      string
	Url      string
	FileName string
	Err      error
	Req      *http.Request
}

//FileContent 存储文件信息
type FileContent struct {
	Content []byte
	Size    int64
	Ext     string
}

var (
	acceptedTypes = []string{
		"jpeg", "bmp", "gif", "png", "jpg",
	}
	infos = []*ImageHandleInfo{
		&ImageHandleInfo{
			Flag:     "/editormd",
			Fromname: "editormd-image-file",
			ToData:   GetRet4Editormd,
		},
		&ImageHandleInfo{
			Flag:     "/ckeditor",
			Fromname: "upload",
			ToData:   GetRet4Ckeditor,
		},
	}
)

func (h *ImageHandle) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if strings.HasPrefix(path, "/image/upload") {
		h.ImageUpload(rw, req)
	} else {
		h.ImageView(rw, req)
	}
}

// ImageUpload 处理图片上传
func (h *ImageHandle) ImageUpload(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	retdata := &RetData{Code: 1, Msg: "ok", Req: req}
	var tinfo *ImageHandleInfo
	for _, info := range infos {
		if strings.Contains(path, info.Flag) {
			tinfo = info
			break
		}
	}
	if tinfo != nil {
		fileinfo, err := h.loadFile(req, tinfo)
		if err != nil {
			retdata.Err = err
		} else {
			filename := strings.Replace(uuid.NewV4().String(), "-", "", -1) + "." + fileinfo.Ext
			datestr := GetCurTime()

			if err1 := saveFile(filepath.Join(h.GetImgDir(), datestr), filename, fileinfo.Content); err == nil {
				retdata.Url = "/image/" + datestr + "/" + filename
			} else {
				retdata.Err = err1
			}
		}
	} else {
		retdata.Err = errors.New("非法访问！")
	}
	rw.Write(tinfo.ToData(retdata))
}

// GetCurTime 得到当前日期，映射到存储路径
func GetCurTime() string {
	now := time.Now()
	bs := []byte("")
	bs = time.Time(now).AppendFormat(bs, "2006/01/02")
	return string(bs)
}

// GetImgDir 文件存储基本目录
func (h *ImageHandle) GetImgDir() string {
	var imgpath = IMG_DIR
	if len(h.imgdir) != 0 {
		imgpath = h.imgdir
	}
	return imgpath //filepath.Join(imgpath, string(bs))
}

// loadFile 得到req中的图片，并判断合法性
func (h *ImageHandle) loadFile(req *http.Request, info *ImageHandleInfo) (*FileContent, error) {
	file, _, err := req.FormFile(info.Fromname)
	if err != nil {
		return nil, err
	}

	fileContent, err := ioutil.ReadAll(file)
	fileMimeTypeStr := http.DetectContentType(fileContent)

	ext, ok := isAccepted(fileMimeTypeStr)
	if !ok {
		return nil, errors.New("非法类型：" + fileMimeTypeStr)
	}
	FileSize := int64(len(fileContent))
	if FileSize > 5*1024*1024 {
		return nil, errors.New("文件过大：")
	}
	return &FileContent{Content: fileContent, Size: FileSize, Ext: ext}, nil
}

// isAccepted 文件类型判断是否合法
func isAccepted(fileMimeTypeStr string) (string, bool) {
	if !strings.HasPrefix(fileMimeTypeStr, "image") {
		return "", false
	}
	for _, acceptedType := range acceptedTypes {
		if strings.HasSuffix(fileMimeTypeStr, "/"+acceptedType) {
			return acceptedType, true
		}
	}
	return "", false
}

// saveFile 保存图片
func saveFile(dir string, filename string, bytes []byte) error {
	os.MkdirAll(filepath.Join(dir), 0777)
	filename = filepath.Join(dir, filename)
	os.Create(filename)
	ioutil.WriteFile(filename, bytes, 0777)
	return nil
}

// GetRet4Editormd Editormd将json字符串转为[]byte
func GetRet4Editormd(data *RetData) []byte {
	var code int64
	msg := "ok"
	if data.Err == nil {
		code = 1
	} else {
		code = 0
		msg = data.Err.Error()
	}
	d := struct {
		Success int64  `json:"success"`
		Message string `json:"message"`
		Url     string `json:"url"`
	}{
		Success: code,
		Message: msg,
		Url:     data.Url,
	}
	databtyes, _ := json.Marshal(d)
	return databtyes
}

// GetRet4Ckeditor Ckeditor将结果转化为[]byte放回
func GetRet4Ckeditor(data *RetData) []byte {
	var (
		url      = ""
		errstr   = ""
		callback = data.Req.FormValue("CKEditorFuncNum")
	)

	if data.Err == nil {
		url = data.Url
	} else {
		errstr = data.Err.Error()
	}
	var d = struct {
		Callback string
		Url      string
		Error    string
	}{
		Callback: callback,
		Url:      url,
		Error:    errstr,
	}

	tpl := `<script type="text/javascript">
			window.parent.CKEDITOR.tools.callFunction({{.Callback}},'{{.Url}}','{{.Error}}');
		</script>`
	tt, _ := template.New("webpage").Parse(tpl)
	bs := &bytes.Buffer{}
	tt.Execute(bs, d)
	return bs.Bytes()
}

// ImageView 查看图片
func (h *ImageHandle) ImageView(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		http.NotFound(w, r)
	} else {
		http.StripPrefix("/image",
			http.FileServer(http.Dir(h.GetImgDir()))).ServeHTTP(w, r)
	}
}
