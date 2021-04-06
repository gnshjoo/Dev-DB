package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gnshjoo/Dev-DB/controllers/db"
	"github.com/gnshjoo/Dev-DB/models"
	"net/http"
	"strconv"
)

// @title GetDBList
// @Description Get DB List from database
// @Accept json
// @Produce json
// @Param Authorization body string true "token"
// @Success 200 {object} models.DataBaseList
// @Router /v1/db/list [get]
func GetDBList(c *gin.Context) {
	var res []models.DataBaseList

	ls, err := db.DBListGet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	for _, l := range ls {
		ok := db.PingDB(l.DBType, l.DBUser, l.DBPass, l.Host, strconv.Itoa(l.Port), l.DatabaseName)
		l.Status = ok

		res = append(res, l)
	}

	c.JSON(http.StatusOK, res)
	return
}

// @title GetDBDetail
// @Description Get DB database list and detail
// @Accept json
// @Produce json
// @Param Authorization body string true "token"
// @Success 200 {object} models.DBDetail
// @Router /v1/db/detail/:id [get]
func GetDBDetail(c *gin.Context) {
	id := c.Param("id")

	res, err := db.DBDeatil(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
