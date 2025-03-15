package controller

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/service"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/utility"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllCourseData(c *gin.Context) {
	data, err := service.GetAllCourseData()
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func GetCourseData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	data, err := service.GetCourseDataLazy(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func GetCourseDataEager(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	data, err := service.GetCourseDataEager(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func PostCourseData(c *gin.Context) {
	var course model.Course
	err := c.ShouldBind(&course)
	if err != nil {
		utility.ClientErrorResponse(c, "Failed to save data. Bad Request body.", err)
		return
	}

	data, err := service.PostCourseData(&course)
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to save data.", err)
		return
	}
	utility.CreatedResponse(c, "Successfully save the data", data)
}

func DeleteCourseData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ServerErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	err = service.DeleteCourseData(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to delete data.", err)
		return
	}
	utility.SuccessResponse(c, "Successfully delete the data", nil)
}
