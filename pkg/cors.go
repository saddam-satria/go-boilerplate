package pkg

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CorsConfig = cors.Config{
	AllowMethods:     []string{"GET", "POST", "UPDATE", "DELETE"},
	AllowFiles:       false,
	AllowAllOrigins:  true,
	MaxAge:           12 * time.Hour,
	AllowCredentials: true,
	AllowHeaders:     []string{"Content-Type", "Origin"},
}
