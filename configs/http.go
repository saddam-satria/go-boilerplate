package configs

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Http struct {
	Engine *gin.Engine
}

func (ctx *Http) Listen(port string, host string) error {

	err := ctx.Engine.Run(fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		return err
	}

	log := Logger{}
	ctx.Engine.Use(log.HttpLogger)

	return nil
}
