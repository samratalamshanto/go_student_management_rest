package controller

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/service"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/utility"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllInvoiceData(c *gin.Context) {
	data, err := service.GetAllInvoiceData()
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func GetInvoiceData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utility.ClientErrorResponse(c, "Id parsing error. Id must be int type.", err)
	}

	data, err := service.GetInvoiceData(uint(id))
	if err != nil {
		utility.ServerErrorResponse(c, utility.CommonFailedToGetMsg, err)
		return
	}
	utility.SuccessResponse(c, utility.CommonSuccessToGetMsg, data)
}

func PostInvoiceData(c *gin.Context) {
	var invoice model.Invoice
	err := c.ShouldBind(&invoice)
	if err != nil {
		utility.ClientErrorResponse(c, "Failed to save data. Bad Request body.", err)
		return
	}

	data, err := service.PostInvoiceData(&invoice)
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to save data.", err)
		return
	}
	utility.CreatedResponse(c, "Successfully save the data", data)
}
