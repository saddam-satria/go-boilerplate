package configs

import "github.com/gin-gonic/gin"

func Init(httpRoute HttpRoute) {
	httpEngine := gin.Default()

	http := Http{
		Engine: httpEngine,
	}

	route := http.Engine.Group("/api/v1")

	httpRoute(route)

	if err := DBConnect(); err != nil {
		panic(err)
	}

	if err := http.Listen("5000", "localhost"); err != nil {
		panic(err)
	}

}
