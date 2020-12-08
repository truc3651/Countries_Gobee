package db

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego/client/orm"
)

func ConnectToPostgres(){
	orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase(
		"default", 
		"postgres", 
		"postgres://postgres:anhyeuhanoi@localhost/VATNOW?sslmode=disable")
	// modified this line
	// postgres://<username>:<password>@localhost/<dbname>?sslmode=disable
	if err != nil {
		log.Fatal("@register: ", err)
	}
}