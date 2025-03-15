package routes

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/controller"
	"github.com/gin-gonic/gin"
	"os"
)

func TeacherRoutes(r *gin.Engine) {
	version, errBool := os.LookupEnv("API_VERSION")

	if errBool {
		version = "/v1"
	}
	routes := r.Group(version + "/api/teacher")
	{
		routes.GET("/", controller.GetAllTeacherData)
		routes.GET("/:id", controller.GetTeacherData)
		routes.POST("/", controller.PostTeacherData)
		routes.DELETE("/:id", controller.DeleteTeacherData)
	}
}
