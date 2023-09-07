package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db sql.DB

func SqlStart() {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/prototype_v0")

	if err != nil {
		panic(err)
	}
	db = *dbt
	dbt.Close()
}

func SqlStop() {
	db.Close()
}
