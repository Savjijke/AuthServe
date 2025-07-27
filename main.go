package main

import (
	"SRC/database"
	"SRC/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	routes.RegisterAuthRoutes(r)

	r.Run(":8080")
}
