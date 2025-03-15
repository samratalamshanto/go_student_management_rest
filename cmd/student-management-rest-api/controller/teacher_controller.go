package controller

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/service"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/utility"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllTeacherData(c *gin.Context) {
	data, err := service.GetAllTeacherData()
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func GetTeacherData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	data, err := service.GetTeacherData(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func GetTeacherDataEager(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	data, err := service.GetTeacherDataEager(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func PostTeacherData(c *gin.Context) {
	var teacher model.Teacher
	err := c.ShouldBind(&teacher)
	if err != nil {
		utility.ClientErrorResponse(c, "Failed to save data. Bad Request body.", err)
		return
	}

	data, err := service.PostTeacherData(&teacher)
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to save data.", err)
		return
	}
	utility.CreatedResponse(c, "Successfully save the data", data)
}

func DeleteTeacherData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	err = service.DeleteTeacherData(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to delete data.", err)
		return
	}
	utility.SuccessResponse(c, "Successfully delete the data", nil)
}
