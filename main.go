package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saddam-satria/go-boilerplate/cmd"
	"github.com/saddam-satria/go-boilerplate/pkg"
)

func main() {

	if err := pkg.Connect(pkg.DbConfig); err != nil {
		panic("Database Failed To Connect" + err.Error())
	}

	args := os.Args

	if len(args) > 1 && args[1] == pkg.MIGRATE_SCRIPT {
		migrate(pkg.Connection)
		return
	}

	r := gin.New()
	r.Use(cors.New(pkg.CorsConfig))
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] %s %s - %v\n",
				params.TimeStamp.Format(time.RFC3339),
				params.Method,
				params.Path,
				params.Latency,
			)
		},
		Output: os.Stdout,
	}))

	ginMode := func() string {
		if pkg.GO_DEBUG == "true" {
			return gin.DebugMode
		}

		return gin.ReleaseMode
	}

	gin.SetMode(ginMode())

	cmd.Routes(r)

	if err := r.Run(":" + pkg.PORT); err != nil {
		panic("Server Failed To Start")
	}

	fmt.Println("Server Running On PORT: " + pkg.PORT)
}
