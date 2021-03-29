package users

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gnshjoo/Dev-DB/DB"
	"github.com/gnshjoo/Dev-DB/controllers/token"
	"log"
	"time"
)

func CreateUser(email, password string) (string, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	db, err := DB.ConnectDB()
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer db.Close()

	query := `insert into users(email, password) values (? , ?)`
	_, err = db.ExecContext(ctx, query, email, CryptoToPw(password))
	if err != nil {
		log.Println(err)
		return "", err
	}
	token, err := token.CreateToken(email)
	if err != nil {
		return "", err
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

