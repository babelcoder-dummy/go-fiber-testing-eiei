package config

import (
	"log"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: Env.DatabaseURL}))
	if err != nil {
		log.Panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Product{}, &model.User{})
	log.Println("Database Migrated")
}
