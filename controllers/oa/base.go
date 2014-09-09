package oa

import (
	"github.com/astaxie/beego"
	"github.com/baiyuxiong/gooa/context"
	"log"
	"strings"
)

type baseController struct {
	beego.Controller
	IsLoggedIn     bool
	moduleName     string
	controllerName string
	actionName     string
}

func (this *baseController) Prepare() {
	log.Println("---------- baseController Prepare called.-----------------")
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "oa"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.Data["context"] = context.AppCtx

	this.IsLoggedIn, _ = this.GetSession("IsLoggedIn").(bool)

	log.Println("IsLoggedIn", this.IsLoggedIn)
	this.Data["IsLoggedIn"] = this.IsLoggedIn
}

func (this *baseController) display(tpl string) {
	this.Layout = "oa/layout.html"
	this.TplNames = this.moduleName + "/" + tpl + ".tpl"

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Breadcrumbs"] = "oa/comm/breadcrumbs.tpl"
}
