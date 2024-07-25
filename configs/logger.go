package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

type Logger struct {
}

func (ctx *Logger) dbLogger() (*logger.Interface, error) {

	fileSys := Fs{
		FileName:   "database.log",
		FolderPath: "./logs",
	}

	file, err := fileSys.GetFile(os.O_CREATE | os.O_WRONLY | os.O_APPEND)

	if err != nil {
		return nil, err
	}

	logger := logger.New(
		log.New(file, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	return &logger, nil
}

func (ctx *Logger) defaultLogger(message string) error {
	fileSys := Fs{
		FileName:   "access.log",
		FolderPath: "./logs",
	}

	file, err := fileSys.GetFile(os.O_CREATE | os.O_WRONLY | os.O_APPEND)

	if err != nil {
		return err
	}

	if _, err := file.Write([]byte(message)); err != nil {
		return err
	}

	return nil
}

func (ctx *Logger) HttpLogger(http *gin.Context) {
	env := Env{
		FileName: ".env",
	}
	message := fmt.Sprintf("[%s] - %s %s %s %s %d", time.Now().Format("2006-01-02 15:04:05"), http.ClientIP(), http.Request.Method, http.Request.URL.Path, http.Request.Proto, http.Writer.Status())
	isDebug := env.GetEnv("DEBUG")

	if isDebug == "false" {
		ctx.defaultLogger(message)
	}

	fmt.Println(message)
}
