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
	// this.TplNames = "home.html"
	this.TplNames = "base/base.html"
}
