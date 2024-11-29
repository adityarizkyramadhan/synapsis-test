package utils

import (
	"github.com/gin-gonic/gin"
)

func ResponseSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}

func ResponseError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"status": "error",
		"error":  err.Error(),
	})
}
