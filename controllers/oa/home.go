package oa

import (
	"github.com/baiyuxiong/gooa/utils"
)

type HomeController struct {
	baseController
}

func (this *HomeController) Get() {
	if !this.IsLoggedIn {
		url := utils.U("")
		this.Redirect(url, 302)
	}
	this.display("home")
}
