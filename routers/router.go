package routers

import (
	"github.com/astaxie/beego"
	"monitor-master/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Receive", &controllers.ReceiveController{})
}
