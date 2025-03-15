package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Err     interface{} `json:"err"`
}

func SuccessResponse(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, CommonResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: msg,
		Msg:     msg,
		Error:   nil,
		Err:     nil,
		Data:    data,
	})
}

func CreatedResponse(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusCreated, CommonResponse{
		Code:    http.StatusCreated,
		Success: true,
		Message: msg,
		Msg:     msg,
		Error:   nil,
		Err:     nil,
		Data:    data,
	})
}

func ClientSideErrorResponse(c *gin.Context, err interface{}, errMsg string) {
	c.JSON(http.StatusBadRequest, CommonResponse{
		Code:    http.StatusBadRequest,
		Success: false,
		Message: errMsg,
		Msg:     errMsg,
		Error:   err,
		Err:     err,
		Data:    nil,
	})
}

func ServerSideErrorResponse(c *gin.Context, err interface{}, errMsg string) {
	c.JSON(http.StatusInternalServerError, CommonResponse{
		Code:    http.StatusInternalServerError,
		Success: false,
		Message: errMsg,
		Msg:     errMsg,
		Error:   err,
		Err:     err,
		Data:    nil,
	})
}
