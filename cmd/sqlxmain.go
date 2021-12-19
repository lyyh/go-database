package main

import (
	"database/pkg/sqlxdemo"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err, db := sqlxdemo.InitDB()
	if err != nil {
		fmt.Println("err:", err)
	}
	sqlxdemo.NameQuery(db)
}
