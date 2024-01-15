package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/saddam-satria/go-boilerplate/cmd/app"
)

func Routes(r *gin.Engine) {
	app.AppRoute(r)
}
