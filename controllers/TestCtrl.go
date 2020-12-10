package controllers

import (
	"Countries-Beego/models"
)

// route "/"
func (c *MainController) GetAllTest() {
	listTest := models.GetAllTests()
	c.Data["json"] = &listTest
    c.ServeJSON()
}