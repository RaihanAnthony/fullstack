package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Raihanpoke/shop-clothing/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/grom/logger"
)

func DBSet() (*gorm.DB, error) {
	dsn := "root:0987@tcp(localhost:3306)/fullstack"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.silent)}) // use the gorm logger package for logging
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// set how mauch min connection
	sqlDB.SqtMaxIdleConns(10)
	// set how long connection will lose
	sqlDB.SetConnMaxIDleTime(10 * time.Minute)
	// set how long can use
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	// Auto migration the model
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("successfuly connected to mysql")
	return db, nil
}
