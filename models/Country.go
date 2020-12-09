package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego/client/orm"
)

// declare Countries struct
type Countries struct {
	Id int 
	CountryCode string
	CountryName string
}

// actions on data

// return all data in table Countries
func GetAll() []Countries {
	o := orm.NewOrm()

	var listCountries []Countries
	qs := o.QueryTable(TableName())
	_, err := qs.All(&listCountries)
	CheckErr("@GetAll: ", err)
	return listCountries
}

// read data in file CSV
// then return two dimensional array of data (code, name)
func ReadCsv(fileName string) [][]string {
	file, err := os.Open(fileName)
	CheckErr("@ReadCsv: ", err)
	records, err := csv.NewReader(file).ReadAll()
	CheckErr("@ReadCsv: ", err)
	return records
}

// with the input args - two dimensional array of data (code, name)
// loop and return array of Countries
func LoopCsv(fileName string) []Countries {
	records := ReadCsv(fileName)
	var listCountries []Countries
	flag := false
	code := -1
	name := -1

	for _, record := range records {
		if !flag {
			if record[0] == "Name" {
				name = 0
				code = 1
			} else {
				name = 1
				code = 0
			}
			flag = true
		}else{
			country := new(Countries)
			country.CountryCode = record[code]
			country.CountryName = record[name]

			listCountries = append(listCountries, *country)
		}
	}
	return listCountries
}

// with the input args - Countries variable
// execute insert into table Countries
func InsertOne(country Countries){
	o := orm.NewOrm()

	_, err := o.Insert(&country)
	CheckErr("@InsertOne: ", err)
}

// with the input args - array of Countries 
// execute insert into table Countries
func InsertMulti(listCountries []Countries){
	o := orm.NewOrm()

	_, err := o.InsertMulti(100, listCountries)
	CheckErr("@InsertMulti: ", err)
}

// execute delete all of rows in table Countries
func DeleteAll(){
	o := orm.NewOrm()

	query := fmt.Sprintf("delete from %s", strings.ToLower(TableName()))
	res, err := o.Raw(query).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("@DeleteAll: postgres row affected nums: ", num)
	}
}

// helper
func TableName() string{
	return "Countries"
}

func CheckErr(mess string, err error){
	if err != nil{
		panic(mess + err.Error())
	}
}

// register
func init() {
	orm.RegisterModel(new(Countries))
}