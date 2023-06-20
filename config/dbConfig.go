package config

import (
	"fmt"
	"log"
	"os"

	"github.com/RianIhsan/goBioskop/database"
	"github.com/RianIhsan/goBioskop/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
  if errEnv := godotenv.Load(); errEnv != nil {
    log.Fatal("Error access .env File!")
  }

  var err error
  dsn := os.Getenv("DSN")
  database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("Connection database failed!")
  } else {
    fmt.Println("Connection Success")
  }
}

func RunMigrate() {
  if err := database.DB.AutoMigrate(&models.Film{}); err != nil {
    log.Fatal("Migratioin failed")
  } else {
    fmt.Println("Migratioin Success")
  }
}
