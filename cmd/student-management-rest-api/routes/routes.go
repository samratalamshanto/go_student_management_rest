package routes

import "github.com/gin-gonic/gin"

func RegisterRoute() *gin.Engine {
	r := gin.Default()

	StudentRoutes(r)
	TeacherRoutes(r)
	CourseRoutes(r)
	InvoiceRoutes(r)

	return r
}
