package routers

import (
	"github.com/astaxie/beego"
	"github.com/codepapa/afmobi_landing/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
