package repository

import (
	"fmt"
	"log"

	"github.com/asaoud2022/todo/app/models"
	"github.com/asaoud2022/todo/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var (
	Task = NewTaskRepository()
	User = NewUserRepository()
)

func Connect(c *config.Config) (*gorm.DB, error) {
	/*
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName, c.SSLMode,
		)

	*/
	//dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})
	return db, err
}
