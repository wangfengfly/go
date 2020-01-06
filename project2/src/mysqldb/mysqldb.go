package mysqldb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dhuser:dhdev123@tcp(10.139.22.181)/xiaohua2nd")
	return db, err
}
