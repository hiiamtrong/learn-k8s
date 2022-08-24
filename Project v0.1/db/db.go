package db

import (
	"fmt"
	"log"

	"github.com/hiiamtrong/todo-simple/config"
	"github.com/hiiamtrong/todo-simple/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Cfg.PGHost, config.Cfg.PGUser, config.Cfg.PGPassword, config.Cfg.PGDB, config.Cfg.PGPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgre connected")
	db.AutoMigrate(&models.Todo{})
	log.Println("Postgre migrated")
	DB = db
}
