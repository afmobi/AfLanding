package main

import (
	"github.com/astaxie/beego"
	"github.com/codepapa/AfLanding/controllers"
)

func main() {

	beego.Router("/", &controllers.MainController{})

	beego.Run()
}
