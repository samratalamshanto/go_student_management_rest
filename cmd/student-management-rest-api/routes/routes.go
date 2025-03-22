package routes

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()

	//first pass through auth middleware
	r.Use(middleware.AuthMiddleware())

	AuthRoutes(r)
	StudentRoutes(r)
	TeacherRoutes(r)
	CourseRoutes(r)
	InvoiceRoutes(r)

	return r
}
