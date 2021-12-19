package sqlxdemo

import (
	"fmt"

	"database/pkg/entity"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() (err error, db *sqlx.DB) {
	dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

func NameQuery(db *sqlx.DB) {
	sql := "SELECT * FROM person WHERE username=:Username"
	rows, err := db.NamedQuery(sql, map[string]interface{}{"Username": "程序员A"})
	if err != nil {
		fmt.Println("db.NamedQuery failed:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u entity.Person
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Printf("user:%#v\n", u)
	}
}
