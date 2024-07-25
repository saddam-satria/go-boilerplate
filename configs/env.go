package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	FileName string
}

func (ctx *Env) GetEnv(key string) string {
	godotenv.Load(ctx.FileName)

	return os.Getenv(key)
}
