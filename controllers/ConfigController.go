package controllers

import "encoding/json"

type ConfigController struct {
	BaseController
}

type Config struct {
	Uuid      string
	Interval  int
	UpdateURL string
	Cpu       interface{}
	Mem       interface{}
	Disk      interface{}
	Network   interface{}
	Processes []string
}

func (this *ConfigController) Get() {
	var config Config
	//todo 解析上报uuid来根据具体情况分配新配置
	json.Marshal(&config)
	this.Data["json"] = config
	this.ServeJSON()
}
