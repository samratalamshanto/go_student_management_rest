package controller

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/dto"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/service"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/utility"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *model.User
	if err := c.ShouldBind(&user); err != nil {
		utility.ClientErrorResponse(c, "Failed to parse user information", err)
		return
	}

	reqBody, err := service.CreateUser(user)
	if err != nil {
		utility.ServerErrorResponse(c, "Failed to parse login information", err)
		return
	}
	utility.CreatedResponse(c, "Successfully Created User", reqBody.UserName)
}

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	if err := c.ShouldBind(&loginDto); err != nil {
		utility.ClientErrorResponse(c, "Failed to parse login information", err)
		return
	}

	responseBody, errLogin := service.Login(loginDto)
	if errLogin != nil {
		utility.ServerErrorResponse(c, "Failed to parse login information", errLogin)
		return
	}

	utility.SuccessResponse(c, "Successfully LoggedIn", responseBody)
}
