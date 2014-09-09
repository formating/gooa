package context

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/baiyuxiong/gooa/utils"
	"log"
)

var AppCtx *AppContext
var APPCache cache.Cache

//APP相关上下文信息，启动时初始化，所有用户看到的信息均一致
type AppContext struct {
	Sitename string
}

//用户相关上下文信息，放在缓存中，不同用户信息不同
type UserContext struct {
	Uid      int64 //用户id
	Cid      int64 //公司id
	Username string
	Email    string
}

type CompanyContext struct {
	Name   string
	Phone  string
	Phone1 string
}

type GroupContext struct {
	Gids []int                 //用户所属分组
	Auth map[string]([]string) //权限列表
}

func init() {
	log.Println("---------- Context init called.-----------------")
	AppCtx = &AppContext{}
	AppCtx.Sitename = beego.AppConfig.String("sitename")

	// cache init
	APPCache, _ = cache.NewCache("memory", `{"interval":3600}`)
}

func SetUserContext(uid int64, uc UserContext) {
	APPCache.Put(utils.CK_UC(uid), uc, 3600) //1小时缓存
}

func GetUserContext(uid int64) UserContext {
	return APPCache.Get(utils.CK_UC(uid)).(UserContext)
}

func DeleteUserContext(uid int64) {
	APPCache.Delete(utils.CK_UC(uid))
}

func IsExistUserContext(uid int64) bool {
	return APPCache.IsExist(utils.CK_UC(uid))
}
