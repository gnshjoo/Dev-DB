package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(user string) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "alskdjfhg")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}