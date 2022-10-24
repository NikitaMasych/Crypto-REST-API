package storage

import (
	"customers/types"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlDB(dbURL string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	launchMigration(db)
	return db
}

func launchMigration(db *gorm.DB) {
	err := db.AutoMigrate(&types.Customer{})
	if err != nil {
		log.Fatal("Can't migrate customer")
	}
	err = db.AutoMigrate(&types.Order{})
	if err != nil {
		log.Fatal("Can't migrate orders")
	}
}
