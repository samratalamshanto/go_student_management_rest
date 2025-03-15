package controller

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/service"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/utility"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllStudentData(c *gin.Context) {
	data, err := service.GetAllStudentData()
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func GetStudentData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	data, err := service.GetStudentData(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func GetStudentDataEager(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	data, err := service.GetStudentDataEager(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func PostStudentData(c *gin.Context) {
	var student model.Student
	err := c.ShouldBind(&student)
	if err != nil {
		utility.ClientErrorResponse(c, "Failed to save data. Bad Request body.", err)
		return
	}

	data, err := service.PostStudentData(&student)
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to save data.", err)
		return
	}
	utility.CreatedResponse(c, "Successfully save the data", data)
}

func DeleteStudentData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	err = service.DeleteStudentData(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to delete data.", err)
		return
	}
	utility.SuccessResponse(c, "Successfully delete the data", nil)
}
