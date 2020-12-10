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
	listCountries := models.GetAllCountries()
	// return data on template
	c.Data["listCountries"] = listCountries
	c.TplName = "countries.tpl"

	// return json
	// c.Data["json"] = &listCountries
    // c.ServeJSON()
}

// route "/insert"
// execute DeleteAll() - delete all of rows in table Countries
// execute LoopCsv() - return array of Countries
// execute InsertMulti(<list>) - insert all of Countries in list, into table Countries
// after inserted, redirect to route "/"
func (c *MainController) InsertCountries() {
	// delete all
	models.DeleteAllCountries()
	// insert
	// uri := "static/files/country.csv"
	uri := "data/country.csv"
	listCountries := models.LoopCountriesCsv(uri)
	models.InsertMultiCountries(listCountries)

	c.Redirect("/", 302)
}

// route "/location"
func (c *MainController) ChooseLocations() {
	listCountries := models.GetAllCountries()
	c.Data["listCountries"] = listCountries
	c.TplName = "locations.tpl"
}


