package users

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gnshjoo/Dev-DB/DB"
	"github.com/gnshjoo/Dev-DB/controllers/token"
	"github.com/gnshjoo/Dev-DB/models"
	"log"
	"time"
)

func CreateUser(email, password string) (*models.TokenDetails, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	db, err := DB.ConnectDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	query := `insert into users(email, password) values (? , ?)`
	_, err = db.ExecContext(ctx, query, email, CryptoToPw(password))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	token, err := jwt.CreateToken(email)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CryptoToPw(pw string) string {
	hash := sha256.New()

	hash.Write([]byte(pw))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}
