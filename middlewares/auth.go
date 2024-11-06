package middlewares

import (
	"github.com/gin-gonic/gin"
	"rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"message": "Authorization token is required"})
	}
	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	context.Set("userId", 1)
	context.Next()
}
