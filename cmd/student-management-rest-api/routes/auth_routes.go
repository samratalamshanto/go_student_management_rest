package routes

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/create-user", controller.CreateUser)
	}
}
