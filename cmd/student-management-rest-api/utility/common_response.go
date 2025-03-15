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

func SuccessResponse(c *gin.Context, msg string, data interface{}) {
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

func CreatedResponse(c *gin.Context, msg string, data interface{}) {
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

func ClientErrorResponse(c *gin.Context, errMsg string, err interface{}) {
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

func ServerErrorResponse(c *gin.Context, errMsg string, err interface{}) {
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
