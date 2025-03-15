package database

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file")
	}

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, conErr := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if conErr != nil {
		return fmt.Errorf("Error connecting to database")
	}

	log.Println("Successfully connected to database")
	DB = db

	//auto migrate
	errMigrate := db.AutoMigrate(&model.Student{}, &model.Teacher{}, &model.Course{}, &model.Invoice{})
	if errMigrate != nil {
		return fmt.Errorf("Error Migrating database")
	}

	return nil
}
