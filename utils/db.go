package utils

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbcon *gorm.DB

func CreateDBconnection() *gorm.DB {
	dsn := ConnectionString

	// checking for existing connections - if true -> close this connection
	if dbcon != nil {
		CloseDBConnection(dbcon)
	}

	// open db connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// setting mysql db settings
	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 15)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	dbcon = db

	log.Println("DB connection was created")
	return db
}

func CloseDBConnection(conn *gorm.DB) {

	db, err := conn.DB()
	if err != nil {
		log.Printf("DB closing is failed: %v", err)
	}
	defer db.Close()
}
