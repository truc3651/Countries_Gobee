package controllers

import (
	"Countries-Beego/models"

	beego "github.com/astaxie/beego/server/web"
)

type MainController struct {
	beego.Controller
}

// route "/"
func (c *MainController) Get() {
	listCountries := models.GetAll()
	c.Data["listCountries"] = listCountries
	c.TplName = "countries.tpl"
}

// route "/insert"
// execute DeleteAll() - delete all of rows in table Countries
// execute LoopCsv() - return array of Countries
// execute InsertMulti(<list>) - insert all of Countries in list, into table Countries
// after inserted, redirect to route "/"
func (c *MainController) InsertCountries() {
	// delete all
	models.DeleteAll()
	// insert
	uri := "static/files/country.csv"
	listCountries := models.LoopCsv(uri)
	models.InsertMulti(listCountries)

	c.Redirect("/", 302)
}


