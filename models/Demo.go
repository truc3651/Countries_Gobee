package models

import "github.com/astaxie/beego/client/orm"

type Demos struct {
	Id int
	CountryCode string 
	CountryName string
}

func GetAllDemos() []Demos {
	o := orm.NewOrm()

	var listDemos []Demos
	qs := o.QueryTable("Demos")
	_, err := qs.All(&listDemos)
	CheckErr("@GetAllDemos: ", err)
	return listDemos
}

func init() {
	orm.RegisterModel(new(Demos))
}