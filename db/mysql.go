package db

import (
	"database/sql"
	"fmt"
)

var mysqlDB *sql.DB

//链接mysql
func InitMysqlDb() *sql.DB {
	if mysqlDB != nil {
		return mysqlDB
	}
	mysqlDB, err := sql.Open("mysql", "root:940805@tcp(localhost:3306)/dev?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	return mysqlDB
}
