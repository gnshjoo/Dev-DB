package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func PingDB(ty, user, pw, host, port, dbName string) bool {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s/%s", user, pw, host, port, dbName)
	db, err := sql.Open(ty, dns)
	if err != nil {
		log.Println(err)

	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return false
	}
	db.Close()
	return true
}
