package pkg

var DbConfig = DatabaseConnection{
	HOST:     GoDotEnv("DB_HOST"),
	USER:     GoDotEnv("DB_USER"),
	PASSWORD: GoDotEnv("DB_PASSWORD"),
	DATABASE: GoDotEnv("DB_DATABASE"),
	PORT:     GoDotEnv("DB_PORT"),
}

var PORT = GoDotEnv("PORT")
var GO_DEBUG = GoDotEnv("DEBUG")
var MIGRATE_SCRIPT = GoDotEnv("MIGRATE_SCRIPT")
