package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	Account string
}

func (user *User) getDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dhuser:test@tcp(10.139.22.181)/testdb")
	return db, err
}

func (user *User) GetById(id int) bool {
	db,err := user.getDB()
	if err != nil {
		fmt.Println("getDB fail."+err.Error())
		return false
	}

	query := fmt.Sprintf("select id,account from user where id=%d", id)
	result,err := db.Query(query)
	if err != nil {
		return false
	}

	result.Next()
	result.Scan(&user.Id, &user.Account)
	return true
}
