package handlers

import (
	"SRC/database"
	"SRC/models"
	"SRC/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось обработать пароль"})
		return
	}

	user := models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}
	err = database.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось создать пользователя"})
		return
	}

	createdUser, err := database.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка после создания пользователя"})
		return
	}

	token, err := utils.GenerateJWT(uint(createdUser.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось создать токен",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Пользователь успешно создан",
		"token":   token,
	})
}

func Login(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.GetUserByEmail(input.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный email или пароль"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный email или пароль"})
		return
	}

	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось создать токен",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Успешный вход!",
		"token":   token,
	})
}
