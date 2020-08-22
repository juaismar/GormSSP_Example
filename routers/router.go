package routers

import (
	"GormSSP_Example/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/simple/paginate", &controllers.MainController{}, "get:SimplePaginate")
}
