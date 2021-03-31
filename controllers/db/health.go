package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gnshjoo/Dev-DB/models"
	"log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"time"
)

func PingDB(ty, user, pw, host, port, dbName string) bool {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pw, host, port, dbName)
	db, err := sql.Open(ty, dns)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return false
	}
	return true
}

func GetDatabase(ty, user, pw, host, port, dbName string) (models.DBDetail, error){
	var dbl models.DBDetail
	var db string
	var dbs []string
	var version string

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pw, host, port, dbName)
	conn, err := sql.Open(ty, dns)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err != nil {
		log.Println(err)
		return dbl, err
	}
	query := `show databases`
	versionQuery := `SELECT VERSION()`
	if ty == "mysql" {
		query = `show databases`
	} else if ty == "postgres" {
		query = `select datname from pg_database`
	}

	rows, err := conn.QueryContext(ctx, query)
	for rows.Next() {
		err := rows.Scan(
			&db,
		)
		if err != nil {
			return dbl, err
		}
		dbs = append(dbs, db)
	}
	if err = rows.Err(); err != nil {
		return dbl, err
	}
	rErr := conn.QueryRowContext(ctx, versionQuery).Scan(&version)
	if rErr != nil {
		return dbl, nil
	}
	dbl.Database = dbs
	dbl.DBVersion = version
	return dbl, nil
}