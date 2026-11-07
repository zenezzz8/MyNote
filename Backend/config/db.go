package config

import (
	"Backend/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnDB() *gorm.DB {
	// read .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .nv")
	}

	//
	dsn := os.Getenv("POSTGRES_URL")

	// open db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("fail connect to database", err)
	}

	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Note{})
}
