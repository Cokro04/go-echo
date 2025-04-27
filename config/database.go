package config

import (
	"database/sql"
	"log"
	"simple-store-app/schemas"

	_ "modernc.org/sqlite" // Import ModernC SQLite driver
)

var DB *sql.DB

func InitDB() {
	var err error
	// Use "sqlite" as the driver name for ModernC
	DB, err = sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize tables
	schemas.InitSchema(DB)

	seedData()
}

func seedData() {
	// Seed users
	_, err := DB.Exec(`INSERT OR IGNORE INTO users (username, password) VALUES
		('admin', 'adminpassword'),
		('user1', 'user1password');`)
	if err != nil {
		log.Fatal(err)
	}

	// Seed stores
	_, err = DB.Exec(`INSERT OR IGNORE INTO stores (name, address) VALUES
		('Store A', '123 Main Street'),
		('Store B', '456 Elm Street');`)
	if err != nil {
		log.Fatal(err)
	}
}
