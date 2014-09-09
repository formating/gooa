package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

func U(s string) string {
	return beego.AppConfig.String("wwwroot") + s
}

func IsLoggedIn() bool {

	return false
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/*Cache key定义 START*/
//用户缓存主键
func CK_UC(uid int64) string {
	return "uc" + fmt.Sprintf("%d", uid)
}

//公司缓存主键
func CK_CC(cid int64) string {
	return "cc" + fmt.Sprintf("%d", cid)
}

//权限缓存主键
func CK_GC(gid int64) string {
	return "gc" + fmt.Sprintf("%d", gid)
}

/*Cache key定义 END*/

/*包装错误信息 START*/
func WrapErrors(errors []*validation.ValidationError) string {
	var errorString string = "<div class=\"alert alert-info\"><button type=\"button\" class=\"close\" data-dismiss=\"alert\">×</button>"
	for _, err := range errors {
		errorString += "<p>" + err.Message + "</p>"
	}
	errorString += "</div>"
	return errorString
}

func WrapString(msg string) string {
	var errorString string = "<div class=\"alert alert-info\"><button type=\"button\" class=\"close\" data-dismiss=\"alert\">×</button>" + msg + "</div>"
	return errorString
}

/*包装错误信息 END*/

//add function for template
//http://beego.me/docs/mvc/view/template.md
func init() {
	beego.AddFuncMap("u", U)
}
