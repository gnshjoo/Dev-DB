package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnshjoo/Dev-DB/handlers"
	"github.com/gnshjoo/Dev-DB/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	_ "github.com/gnshjoo/Dev-DB/docs"
)

// @title Swagger Dev-DB
// @version 0.0.1
// @BasePath /v1
func main() {
	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// routes
	v1 := r.Group("/v1")
	{
		v1.POST("/users/login", handlers.UserLogin)
		v1.POST("/users/logout", middleware.TokenAuthMiddleware(), handlers.UserLogout)
		v1.POST("/users/signup", handlers.UserSignUp)
		//v1.POST("/token/refresh", users.Refresh)
		v1.GET("/db/list", middleware.TokenAuthMiddleware(), handlers.GetDBList)
		v1.GET("/db/detail/:id", middleware.TokenAuthMiddleware(), handlers.GetDBDetail)
	}

	log.Fatal(r.Run(":8080"))
}
