package controllers

import "github.com/astaxie/beego"

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["json"] = map[string]string{
		"name":    beego.BConfig.AppName,
		"version": beego.AppConfig.String("version"),
		"runmode": beego.BConfig.RunMode,
	}
	this.ServeJSON()
}
