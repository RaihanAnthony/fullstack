package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConenection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "0987"
	lcHost := "localhost:3306"
	dbName := "fullstack"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+lcHost+")/"+dbName)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	// make sure connection to mysql
	err = db.Ping()
	if err != nil {
		log.Println("failed connection to mysql", err)
		return nil, err
	}

	return db, nil
}
