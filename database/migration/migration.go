package migration

import (
	"auth-jwt/database"
	"auth-jwt/models/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrate")
}
