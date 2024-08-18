package dtos

import "github.com/gin-gonic/gin"

func SuccessResponse(data interface{}, message string) gin.H {
	return gin.H{
		"data":    data,
		"message": message,
	}
}

func ErrorResponse(message string) gin.H {
	return gin.H{
		"data":    nil,
		"message": message,
	}
}
