package db

import (
	"context"
	"github.com/gnshjoo/Dev-DB/DB"
	"github.com/gnshjoo/Dev-DB/models"
	"strconv"
	"time"
)

func DBListGet() ([]models.DataBaseList, error) {
	var ls []models.DataBaseList

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	conn, err := DB.ConnectDB()
	if err != nil {
		return nil, err
	}

	query := `select id, db_name, db_user, db_password, db_type, port, host from databaselist`
	rows, err := conn.QueryContext(ctx, query)

	defer rows.Close()

	for rows.Next() {
		var l models.DataBaseList

		err := rows.Scan(
			&l.ID,
			&l.DatabaseName,
			&l.DBUser,
			&l.DBPass,
			&l.DBType,
			&l.Port,
			&l.Host,
		)

		if err != nil {
			return ls, err
		}
		ls = append(ls, l)
	}
	if err = rows.Err(); err != nil {
		return ls, err
	}
	return ls, nil
}

func DBDeatil(id string) (models.DBDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var u models.DataBaseList
	var db models.DBDetail

	conn, err := DB.ConnectDB()
	defer conn.Close()
	if err != nil {
		return db, err
	}
	ll, err := strconv.Atoi(id)
	if err != nil {
		return db, err
	}

	query := `select id, db_name, db_user, db_password, port, db_type, host from databaselist where id=?`

	row := conn.QueryRowContext(ctx, query, ll)
	rErr := row.Scan(
		&u.ID,
		&u.DatabaseName,
		&u.DBUser,
		&u.DBPass,
		&u.Port,
		&u.DBType,
		&u.Host,
	)
	if rErr != nil {
		return db, err
	}
	res, err := GetDatabase(u.DBType, u.DBUser, u.DBPass, u.Host, strconv.Itoa(u.Port), u.DatabaseName)
	if err != nil {
		return db, err
	}

	return res, nil
}