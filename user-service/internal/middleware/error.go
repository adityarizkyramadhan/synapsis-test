package middleware

import (
	"github.com/adityarizkyramadhan/synapsis-test/user-service/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		if err != nil {
			errParse := utils.ParseError(err.Error())
			utils.ErrorResponse(ctx, errParse.StatusCode, errParse.Message)
		}
	}
}
