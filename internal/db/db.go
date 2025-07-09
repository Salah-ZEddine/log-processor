package db

import (
	"fmt"
	"log"
	"os"
	"github.com/Salah-ZEddine/incident-dashboard-common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//setup the database connection and auto-migrate tables and columns
func Init(autoMigrate bool){

	//Format the connection string (DSN)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	//open the connection to the database using GORM and Postgres Sql driver
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	if autoMigrate {
		err = DB.AutoMigrate(&models.Log{})
		if err != nil {
		log.Fatal("Failed to automigrate tables: ", err)
		}
	}

	log.Println("Database connected")
	
}

func SaveLog(logEntry *models.Log) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}
	return DB.Create(logEntry).Error
}