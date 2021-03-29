package main

import (
	"github.com/gin-gonic/gin"
	users "github.com/gnshjoo/Dev-DB/handlers"
	"github.com/gnshjoo/Dev-DB/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @title Swagger Dev-DB
// @version 0.0.1
// @BasePath /v1

func main() {
	r := gin.New()

	r.Use(gin.Recovery())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// routes
	v1 := r.Group("/v1")
	{
		v1.POST("/users/login", users.UserLogin)
		v1.POST("/users/logout", middleware.TokenAuthMiddleware(), users.UserLogout)
		v1.POST("/users/signup", users.UserSignUp)
		//v1.POST("/token/refresh", users.Refresh)
	}

	log.Fatal(r.Run(":8080"))
}
