package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize database connection
func initDB() {
	var err error
	db, err = sql.Open("mysql", "your-db-connection-string")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")
}

// Close database connection
func closeDB() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Closed database connection")
}
