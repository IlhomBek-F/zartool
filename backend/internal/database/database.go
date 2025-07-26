package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	DB  gorm.DB
	Env Config
}

func App() Application {
	config := NewConfig()

	te := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		config.host, config.username, config.password, config.database, config.DbPort, config.schema)

	db, err := gorm.Open(postgres.Open(te), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to db")
	}

	return Application{DB: *db, Env: config}
}
