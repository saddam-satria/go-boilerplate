package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/saddam-satria/go-boilerplate/configs"
)

func Init(httpRoute *gin.RouterGroup) {
	httpRoute.GET("student", func(ctx *gin.Context) {
		handler := api.ResponseHandler{
			Engine: ctx,
		}

		handler.Success("student", http.StatusCreated, nil)
	})
}
