package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	// "strconv"
)

func main() {
	db, err := sql.Open("mysql", "trb:123@tcp(172.16.0.200:3306)/hotel?charset=utf8")
	if err != nil {
		panic(err.Error())
		println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err.Error())
		println(err.Error())
		return
	}
	defer rows.Close()
	var name string
	var phone string
	var passwd string
	var idnum string
	var role string
	var id int
	for rows.Next() {
		err := rows.Scan(&name, &phone, &passwd, &idnum, &role, &id)
		if err != nil {
			panic(err.Error())
			println(err.Error())
			return
		}
		println(name)
	}
}
