package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "afmobi.com"
	this.Data["Email"] = "service@afmobi.com"
	this.TplName = "base/base.html"
}
