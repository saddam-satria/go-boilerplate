package app

import "github.com/gin-gonic/gin"

func AppRoute(c *gin.Engine) {
	c.GET("/", AppHandler)
}
