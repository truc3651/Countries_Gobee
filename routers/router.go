package routers

import (
	"Countries-Beego/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/insert", &controllers.MainController{}, "get:InsertCountries")
}
