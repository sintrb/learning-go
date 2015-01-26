package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Post() {
	requestbody := this.Ctx.Input.CopyBody()
	println("begin")
	println(string(requestbody))
	println("end")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
