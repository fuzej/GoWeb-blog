package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open(`mysql`, `root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true`)
	if err != nil {
		panic(err.Error())
	}

}
