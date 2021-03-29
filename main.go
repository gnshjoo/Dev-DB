package main

import (
	"github.com/gnshjoo/Dev-DB/handlers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
		v1.POST("/users/login", handlers.UserLogin)
		v1.POST("/users/logout", handlers.UserLogout)
		v1.POST("/users/signup", handlers.SignUp)
	}

	r.Run(":8080")
}