package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

/*
func MutiStrutsToStr(arrs []interface{}) string {
	length := len(arrs)
	if length == 0 {
		return ""
	}
	strs := make([]string, 0)
	for i := 0; i < length; i++ {
		strs = append(strs, string(json.Marshal(arrs[i])))
	}
	return string(json.Marshal(strs))
}

func StrToMutiStruct(str string) []interface{} {
	str := strings.TrimSpace(str)
	if len(str) == 0 {
		return nil
	}
	var strs []string
	json.Unmarshal([]byte(str), &strs)
	objs := make([]interface{}, 0)
	for i := 0; i < len(strs); i++ {
		var v interface{}
		json.Unmarshal([]byte(strs[i]), &v)
		objs := append(objs, v)
	}
	return objs
}
*/

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}
