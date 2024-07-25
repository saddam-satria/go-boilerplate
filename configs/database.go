package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

type DatabaseConnection struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
	Ssl      string
	DB       *gorm.DB
}

func (ctx *DatabaseConnection) GetDSN() string {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname= %s port=%s sslmode=%s", ctx.Host, ctx.User, ctx.Password, ctx.Database, ctx.Port, ctx.Ssl)

	return dsn
}

func DBConnect() error {

	loggerConf := Logger{}

	logger, err := loggerConf.dbLogger()

	if err != nil {
		return err
	}

	env := Env{
		FileName: ".env",
	}

	ctx := DatabaseConnection{
		Host:     env.GetEnv("DB_HOST"),
		User:     env.GetEnv("DB_USER"),
		Password: env.GetEnv("DB_DATABASE"),
		Database: env.GetEnv("DB_DATABASE"),
		Port:     env.GetEnv("DB_PORT"),
		Ssl:      env.GetEnv("DB_SSL"),
	}

	db, err := gorm.Open(postgres.Open(ctx.GetDSN()), &gorm.Config{Logger: *logger})

	if err != nil {
		return err
	}

	ctx.DB = db
	Database = db

	return nil
}
