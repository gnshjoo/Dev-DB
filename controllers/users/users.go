package users

import (
	"context"
	"github.com/gnshjoo/Dev-DB/DB"
	"github.com/gnshjoo/Dev-DB/models"
	"log"
	"time"
)

func LoginUser(email, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := DB.ConnectDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var u models.User

	query := `select id, email, password from users where email=? and password=?`

	row := conn.QueryRowContext(ctx, query, email, CryptoToPw(password))
	err = row.Scan(
		&u.ID,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		return nil, err
	}
	if u.Email != "" {
		return &u, nil
	} else {
		return nil, nil
	}
}

func LogoutUser(au string) (bool, error) {
	del, err := DB.DeleteAuth(au)
	if err != nil || del == 0 {
		return false, err
	}
	return true, nil
}
