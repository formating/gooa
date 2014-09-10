package routers

import (
	"github.com/astaxie/beego"
	"github.com/baiyuxiong/gooa/controllers/oa"
	"log"
)

func init() {
	log.Println("---------- routers init called.-----------------")

	beego.Router("/", &oa.MainController{}, "get:Login;post:DoLogin")
	beego.Router("/reg", &oa.MainController{}, "get:Reg;post:DoReg")
	beego.Router("/home", &oa.HomeController{})
	beego.Router("/logout", &oa.MainController{}, "get:Logout;post:Logout")
}
