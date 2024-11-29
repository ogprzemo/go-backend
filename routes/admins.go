package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/models"
	"strconv"
)

func CreateAdmin(context *gin.Context) {
	var admin models.Admin
	err := context.ShouldBindJSON(&admin)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = admin.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error creating admin)"})
		return
	}

	context.JSON(201, admin)
}

func UpdateAdmin(context *gin.Context) {
	_, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse admin id"})
		return
	}
}

func deleteAdmin(context *gin.Context) {
	_, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse admin id"})
		return
	}
}

func GetAllAdmins(context *gin.Context) {
	admins, err := models.GetAllAdmins()
	if err != nil {
		context.JSON(500, gin.H{"error": "Error getting admins"})
		return
	}
	context.JSON(200, admins)
}
