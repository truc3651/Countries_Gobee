package models

import "github.com/astaxie/beego/client/orm"

type Tests struct {
	CountryCode string 
	CountryName string
}

func GetAllTests() []Tests {
	o := orm.NewOrm()

	var listTests []Tests
	qs := o.QueryTable("Tests")
	_, err := qs.All(&listTests)
	CheckErr("@GetAll: ", err)
	return listTests
}

func init() {
	orm.RegisterModel(new(Tests))
}