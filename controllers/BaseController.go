package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	fmt.Println("base")
}
