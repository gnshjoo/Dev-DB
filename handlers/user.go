package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnshjoo/Dev-DB/DB"
	jwt "github.com/gnshjoo/Dev-DB/controllers/token"
	"github.com/gnshjoo/Dev-DB/controllers/users"
	"github.com/gnshjoo/Dev-DB/models"
	jw "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
)

func UserLogin(c *gin.Context) {
	var u *models.User
	err := c.Bind(&u)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	ok, err := users.LoginUser(u.Email, u.Password)
	if ok != nil {
		tk, err := jwt.CreateToken(uint64(ok.ID))
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		saveErr := DB.CreateAuth(uint64(u.ID), tk)
		if saveErr != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		token := map[string]string{
			"access_token":  tk.AccessToken,
			"refresh_token": tk.RefreshToken,
		}
		c.JSON(http.StatusOK, token)
		return
	}
	c.JSON(http.StatusUnauthorized, err)
	return
}

func UserLogout(c *gin.Context) {
	au, err := jwt.ExtractTokenMetadata(c.Request)
	if err != nil || au == nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	ok, err := users.LogoutUser(au.AccessUuid)
	if err != nil || !ok {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Logout!")
	return
}

func UserSignUp(c *gin.Context) {
	var u *models.User
	err := c.Bind(&u)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	tk, err := users.CreateUser(u.Email, u.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	saveErr := DB.CreateAuth(uint64(u.ID), tk)
	if saveErr != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	token := map[string]string{
		"access_token":  tk.AccessToken,
		"refresh_token": tk.RefreshToken,
	}
	c.JSON(http.StatusOK, token)
	return
}

func Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	os.Setenv("REFRESH_SECRET", "asdlfkjsflkjsdfhsdf")

	token, err := jw.Parse(refreshToken, func(token *jw.Token)(interface{}, error) {
		if _, ok := token.Method.(*jw.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}

	if _, ok := token.Claims.(jw.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	claims, ok := token.Claims.(jw.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		del, delErr := DB.DeleteAuth(refreshUuid)
		if delErr != nil || del == 0 {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		ts, createErr := jwt.CreateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		saveErr := DB.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token": ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusOK, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}
