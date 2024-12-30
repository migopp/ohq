package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Spawn() {
	// Open the DB
	//
	// In SQLite it's just a file!
	odb, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to `%s`: %v", name, err)
	}
	db = odb
	fmt.Printf("Connection to `%s` established\n", name)

	// Perform migrations automatically
	//
	// This is based on the schema changes in `models.go`
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("Failed to migrate model `User`")
	}
}

func Close() {
	sdb, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close DB connection")
	}
	if err := sdb.Close(); err != nil {
		log.Fatal("Failed to close DB connection")
	}
	log.Println("DB connection closed")
}
