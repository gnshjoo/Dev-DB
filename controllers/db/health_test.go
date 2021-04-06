package db

import (
	"github.com/gnshjoo/Dev-DB/models"
	"gopkg.in/stretchr/testify.v1/assert"
	"strconv"
	"testing"
)

func TestPingDBError(t *testing.T) {
	v := models.DataBaseList{
		DBType: "mysql",
		DBUser: "gnshjoo",
		DBPass: "gnshjoo12",
		Host: "localhost",
		Port: 1234,
		DatabaseName: "db",
	}

	err :=PingDB(v.DBType, v.DBUser, v.DBPass, v.Host, strconv.Itoa(v.Port), v.DatabaseName)

	assert.False(t, err)
	assert.Equal(t, err, false)
}

func TestGetDatabaseError(t *testing.T) {
	v := models.DataBaseList{
		DBType: "mysql",
		DBUser: "gnshjoo",
		DBPass: "gnshjoo12",
		Host: "192.168.1.104",
		Port: 1234,
		DatabaseName: "db",
	}

	db, err := GetDatabase(v.DBType, v.DBUser, v.DBPass, v.Host, strconv.Itoa(v.Port), v.DatabaseName)

	assert.NotEmpty(t, err)
	assert.Empty(t, db)
	assert.Equal(t, err, "err")
}