package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/models"
	"rest-api/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	context.JSON(201, user)
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	valid, err := user.ValidateCredentials()
	if err != nil {
		context.JSON(500, gin.H{"error": "Error validating credentials"})
		return
	}

	if !valid {
		context.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		context.JSON(500, gin.H{"error": "Error generating token"})
	}

	context.Header("Authorization", token)

	utils.GenerateToken(user.Email, int64(user.ID))

	context.JSON(200, gin.H{"message": "Login successful"})
}
