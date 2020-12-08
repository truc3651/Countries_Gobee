package routers

import (
	"Countries-Beego/controllers"
	"Countries-Beego/db"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	// connect to postgres db
	db.ConnectToPostgres()

	// enums all routers which use for this app
	beego.Router("/", &controllers.MainController{})
	beego.Router("/insert", &controllers.MainController{}, "get:InsertCountries")
}
