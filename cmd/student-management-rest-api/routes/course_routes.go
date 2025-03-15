package routes

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/controller"
	"github.com/gin-gonic/gin"
	"os"
)

func CourseRoutes(r *gin.Engine) {
	version, errBool := os.LookupEnv("API_VERSION")

	if errBool {
		version = "/v1"
	}
	routes := r.Group(version + "/api/course")
	{
		routes.GET("/", controller.GetAllCourseData)
		routes.GET("/:id", controller.GetCourseData)
		routes.POST("/", controller.PostCourseData)
		routes.DELETE("/:id", controller.DeleteCourseData)
	}
}
