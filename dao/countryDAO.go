package dao

import (
	"Countries-Beego/models"
	"fmt"
	"log"

	"github.com/astaxie/beego/client/orm"
)

// TblName, _ := beego.AppConfig.String("TblName")

func getAll() {
	o := orm.NewOrm()
	var countries []*models.Country

	qs := o.QueryTable("TblName")
	num, err := qs.Filter("Name", "nguyen").All(&countries)
	if num>0 && err==nil{
		fmt.Println("@relationSQL: ", countries[0])
	}

}

// connect
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", "root:anhyeuhanoi@/vatnow?charset=utf8")
	if err != nil {
		log.Fatal("@register: ", err)
	}
}

