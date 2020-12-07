package models

import "github.com/astaxie/beego/client/orm"

type Countries struct {
	Id int 
	CountryCode string
	CountryName string
}

func init() {
	orm.RegisterModel(new(Countries))
}