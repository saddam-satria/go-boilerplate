package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(httpRoute *gin.RouterGroup) {
	httpRoute.GET("auth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "auth"})
	})
}
