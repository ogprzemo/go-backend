package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	authenticated.POST("/admin", CreateAdmin)
	authenticated.PUT("/admin", UpdateAdmin)
	authenticated.DELETE("/admin", deleteAdmin)
	authenticated.GET("/admin", GetAllAdmins)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
