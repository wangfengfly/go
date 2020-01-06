package model

import (
	"fmt"
	"mysqldb"
)

type User struct {
	Id int
	Account string
}


func (user *User) GetById(id int) bool {
	db,err := mysqldb.GetDB()
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
