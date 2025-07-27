package routes

import (
	"SRC/handlers"
	"SRC/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}
	JwtGroup := router.Group("/")
	JwtGroup.Use(middleware.AuthMiddleware())
	JwtGroup.GET("/profile", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Добро пожаловать в профиль"})
	})

}
