package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gnshjoo/Dev-DB/DB"
	jwt "github.com/gnshjoo/Dev-DB/controllers/token"
	"github.com/gnshjoo/Dev-DB/controllers/users"
	"github.com/gnshjoo/Dev-DB/models"
	"net/http"
)

func UserLogin(c *gin.Context) {
	var u *models.User
	err := c.Bind(&u)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}
	ok, err := users.LoginUser(u.Email, u.Password)
	if ok {
		tk, err := jwt.CreateToken(u.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
		}
		saveErr := DB.CreateAuth(u.Email, tk)
		if saveErr != nil {
			c.JSON(http.StatusUnauthorized, err)
		}
		token := map[string]string{
			"access_token":  tk.AccessToken,
			"refresh_token": tk.RefreshToken,
		}
		c.JSON(http.StatusOK, token)
	}
	c.JSON(http.StatusUnauthorized, err)
}

func UserLogout(c *gin.Context) {
	au, err := jwt.ExtractTokenMetadata(c.Request)
	if err != nil {
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
	saveErr := DB.CreateAuth(u.Email, tk)
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
	// todo : refresh token
}
