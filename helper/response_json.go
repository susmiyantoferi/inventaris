package helper

import (
	"inventaris/web"

	"github.com/gin-gonic/gin"
)

func ResponseJSON(ctx *gin.Context, code int, status string, data interface{}) {
	ctx.JSON(code, web.WebResponse{
		Code: code,
		Status: status,
		Data: data,
	})
}