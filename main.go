package main

import (
	"afmobi_landing/controllers"
	"github.com/astaxie/beego"
)

func main() {

	beego.Router("/", &controllers.MainController{})

	beego.Run()
}
