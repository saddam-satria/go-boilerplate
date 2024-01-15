package main

import (
	"fmt"

	"github.com/saddam-satria/go-boilerplate/cmd/app"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&app.UserCredential{}); err != nil {
		panic("Migration Failed Reason: " + err.Error())
	}

	fmt.Println("Migrate Success")
}
