package repository

import (
	"fmt"
	"log"

	"github.com/Iagobarros211256/voluryashop/configs"
	"github.com/Iagobarros211256/voluryashop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DDB *gorm.DB

func ConnDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.GetEnv("DB_HOST"),
		configs.GetEnv("DB_USER"),
		configs.GetEnv("DB_PASSWORD"),
		configs.GetEnv("DB_NAME"),
		configs.GetEnv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("faliure. cant connect to database", err)
	}

	DB = db
	db.AutoMigrate(&models.User{})
}
