package controllers

import (
	"Countries-Beego/models"
)

// route "/"
func (c *MainController) GetAllDemos() {
	listDemos := models.GetAllDemos()
	c.Data["json"] = &listDemos
    c.ServeJSON()
}