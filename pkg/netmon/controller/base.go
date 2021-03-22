package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code	int			`json:"code"`
	Message	string		`json:"message"`
	Data	interface{}	`json:"data"`
}

func SuccessResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Message: "OK",
		Data: data,
	})
}

func ErrorResp(c *gin.Context, code, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: code,
		Message: msg,
		Data: data,
	})
}
