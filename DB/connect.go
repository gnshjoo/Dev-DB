package DB

import (
	"database/sql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "gnshjoo:gnshjoo12@tcp(localhost:3306)/DB")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
