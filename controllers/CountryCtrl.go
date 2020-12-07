package controllers

import (
	"Countries-Beego/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	beego "github.com/astaxie/beego/server/web"
	_ "github.com/lib/pq"

	"github.com/astaxie/beego/client/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	listCountries := getAll()
	c.Data["listCountries"] = listCountries
	c.TplName = "countries.tpl"
}

func (c *MainController) InsertCountries() {
	// delete all
	deleteAll()
	// insert
	uri := "C:/Users/USER/go/src/github.com/truc3651/lam cho VATNOW/Countries-Beego/static/files/country.csv"
	listCountries := loopCsv(uri)
	insertMulti(listCountries)

	c.Redirect("/", 302)
}

// DAO
func checkErr(mess string, err error){
	if err != nil{
		panic(mess + err.Error())
	}
}

func getAll() []models.Countries {
	o := orm.NewOrm()

	var listCountries []models.Countries
	qs := o.QueryTable("Countries")
	_, err := qs.All(&listCountries)
	checkErr("@getAll: ", err)
	return listCountries
}

func readCsv(fileName string) [][]string {
	file, err := os.Open(fileName)
	checkErr("@readCsv: ", err)
	records, err := csv.NewReader(file).ReadAll()
	checkErr("@readCsv: ", err)
	return records
}

func loopCsv(fileName string) []models.Countries {
	records := readCsv(fileName)
	var listCountries []models.Countries
	flag := false
	code := 0
	name := 1

	for _, record := range records {
		if !flag {
			if record[0] == "Name"{
				code = 0
				name = 1
			}
			flag = true
		}else{
			country := new(models.Countries)
			country.CountryCode = record[code]
			country.CountryName = record[name]

			listCountries = append(listCountries, *country)
		}
	}
	return listCountries
}

func insertOne(country models.Countries){
	o := orm.NewOrm()

	_, err := o.Insert(&country)
	checkErr("@insertOne: ", err)
}

func insertMulti(listCountries []models.Countries){
	o := orm.NewOrm()

	_, err := o.InsertMulti(100, listCountries)
	checkErr("@insertMulti: ", err)
}

func deleteAll(){
	o := orm.NewOrm()

	res, err := o.Raw("delete from countries").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("@deleteAll: postgres row affected nums: ", num)
	}
}

// connect
func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase(
		"default", 
		"postgres", 
		"postgres://postgres:anhyeuhanoi@localhost/VATNOW?sslmode=disable")
	if err != nil {
		log.Fatal("@register: ", err)
	}
}
