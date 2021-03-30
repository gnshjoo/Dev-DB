package DB

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "gnshjoo:gnshjoo12@tcp(localhost:3306)/DB")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
