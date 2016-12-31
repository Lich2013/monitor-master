package controllers

type ReceiveController struct {
	BaseController
}

func (this *ReceiveController) Post() {
	//todo 接收存储上报信息
	this.Data["json"] = map[string]string{"status": "success"}
}
