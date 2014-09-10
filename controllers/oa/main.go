package oa

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/baiyuxiong/gooa/context"
	"github.com/baiyuxiong/gooa/models"
	"github.com/baiyuxiong/gooa/models/bcompany"
	"github.com/baiyuxiong/gooa/models/buser"
	"github.com/baiyuxiong/gooa/utils"
	"log"
	"strings"
	"time"
)

type MainController struct {
	baseController
}

func (this *MainController) Login() {
	log.Println("MainController IsLoggedIn", this.IsLoggedIn)
	if this.IsLoggedIn {
		url := utils.U("home")
		this.Redirect(url, 302)
	} else {
		flash := beego.ReadFromRequest(&this.Controller)
		this.Data["flash"] = flash.Data
		this.TplNames = "oa/index.tpl"
	}
}
func (this *MainController) DoLogin() {
	var (
		uemail string = strings.TrimSpace(this.GetString("uemail"))
		upw    string = strings.TrimSpace(this.GetString("upw"))
	)

	user, err := buser.FindUserByEmail(uemail)
	log.Println("user " + user.Username)
	log.Println("user.Password " + user.Password)
	log.Println("utils.Md5(upw) " + utils.Md5(upw))
	if (err != nil) || (user.Password != utils.Md5(upw)) || (user.Uid < 1) {
		flash := beego.NewFlash()
		flash.Set("uemail", uemail)
		flash.Set("error", utils.WrapString("登录出错，请检查用户名密码是否正确"))
		flash.Store(&this.Controller)
		url := utils.U("")
		this.Redirect(url, 302)
	} else {
		this.SetSession("IsLoggedIn", true)
		uc := context.UserContext{}
		context.SetUserContext(user.Uid, uc)
		url := utils.U("home")
		this.Redirect(url, 302)
	}

}

func (this *MainController) Reg() {
	flash := beego.ReadFromRequest(&this.Controller)
	this.Data["flash"] = flash.Data
	this.TplNames = "oa/reg.tpl"
}
func (this *MainController) DoReg() {
	var (
		name       string = strings.TrimSpace(this.GetString("name"))
		phone      string = this.GetString("phone")
		address    string = strings.TrimSpace(this.GetString("address"))
		email      string = strings.TrimSpace(this.GetString("email"))
		pw         string = strings.TrimSpace(this.GetString("pw"))
		pw_confirm string = strings.TrimSpace(this.GetString("pw_confirm"))
	)
	flash := beego.NewFlash()

	c := &models.BCompany{Name: name, Phone: phone, Address: address}
	u := &models.BUser{Username: email, Password: pw, Email: email}

	v := validation.Validation{}
	v.Required(c.Name, "name").Message("单位名称不能为空")
	v.Required(c.Phone, "phone").Message("电话不能为空")
	v.Required(c.Address, "address").Message("地址不能为空")
	v.Required(u.Username, "username").Message("邮箱不能为空")
	v.Required(u.Password, "password").Message("密码不能为空")
	v.Email(u.Username, "username").Message("邮箱格式不正确")

	if pw != pw_confirm {
		v.SetError("password", "密码与确认密码不匹配")
	}

	errorString := ""
	if v.HasErrors() {
		errorString = utils.WrapErrors(v.Errors)
	}

	if errorString != "" {
		// validation does not pass
		log.Println(errorString)
		flash.Set("error", errorString)
		flash.Set("name", name)
		flash.Set("phone", phone)
		flash.Set("address", address)
		flash.Set("email", email)
		flash.Store(&this.Controller)

		url := utils.U("reg")
		this.Redirect(url, 302)
	} else {
		//insert into db
		//check email exists
		if buser.IsUserEmailExists(email) {
			v.SetError("username", "邮件已存在，不能重复注册。")
			flash.Set("error", utils.WrapErrors(v.Errors))
			flash.Store(&this.Controller)
			url := utils.U("reg")
			this.Redirect(url, 302)
		} else {
			c.Createdate = time.Now().Unix()
			u.Created_on = time.Now().Unix()

			u.Password = utils.Md5(u.Password)
			cid, err := bcompany.Insert(c)

			if err == nil && cid > 0 {
				u.Cid = cid
				uid, uerr := buser.Insert(u)
				if uerr == nil && uid > 0 {
					url := utils.U("home")
					this.Redirect(url, 302)
				}
			}

			flash.Set("error", utils.WrapString("注册出错，请联系管理员"))
			flash.Store(&this.Controller)
			url := utils.U("reg")
			this.Redirect(url, 302)
		}

	}
}
func (this *MainController) Logout() {
	this.SetSession("IsLoggedIn", false)
	url := utils.U("")
	this.Redirect(url, 302)
}
