package routes

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/controller"
	"github.com/gin-gonic/gin"
	"os"
)

func StudentRoutes(r *gin.Engine) {
	version, errBool := os.LookupEnv("API_VERSION")
	if errBool {
		version = "/v1"
	}

	studentRoutes := r.Group(version + "/api/student")
	{
		studentRoutes.GET("/", controller.GetAllStudentData)
		studentRoutes.GET("/:id", controller.GetStudentData)
		studentRoutes.GET("/eager/:id", controller.GetStudentDataEager)
		studentRoutes.POST("/", controller.PostStudentData)
		studentRoutes.DELETE("/:id", controller.DeleteStudentData)
	}
}
