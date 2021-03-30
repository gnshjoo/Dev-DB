package db

import (
	"context"
	"github.com/gnshjoo/Dev-DB/DB"
	"github.com/gnshjoo/Dev-DB/models"
	"time"
)

func GetDBList() ([]models.DataBaseList, error) {
	var ls []models.DataBaseList

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	conn, err := DB.ConnectDB()
	if err != nil {
		return nil, err
	}

	query := `select id, db_name, db_user, db_password, db_type, port from databaselist`
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