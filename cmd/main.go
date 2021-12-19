package main

import (
	"database/pkg/transaction"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	DB = database
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
}

func main() {
	defer DB.Close()

	// // 实例化一个person
	// p := &entity.Person{
	// 	Username: "程序员A",
	// 	Sex:      "man",
	// 	Email:    "374659635@qq.com",
	// }
	// // 插入一个数据
	// _, err := sqldb.Insert(DB, p)
	// if err != nil {
	// 	fmt.Println("sqldb.Insert err", err)
	// }

	// _, err = sqldb.Delete(DB, "stu001")
	// if err != nil {
	// 	fmt.Println("sqldb.Delete err", err)
	// }

	transaction.TracInsert(DB)
}
