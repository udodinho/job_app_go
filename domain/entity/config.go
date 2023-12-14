package entity

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Job{}) 

	return err
}


func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	configFile := &Config{ 
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),

	}
	
	db, err := Connect(configFile)

	if err != nil {
		log.Fatal("Could not connect to db, ", err)
	}

	err = MigrateDB(db)

	if err != nil {
		log.Fatal("Could not migrate database, ", err)
	}

	DB = db

	fmt.Println("Database connected successfully")
	
}