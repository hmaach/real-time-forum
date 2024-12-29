package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"forum/server/config"
)

var (
	database     = "sqlite3"
	databasePath = config.BasePath + "server/database/database.db"
	schemaPath   = config.BasePath + "server/database/schema.sql"
	seedPath     = config.BasePath + "server/database/seed.sql"
	DB           *sql.DB
)

func Connect() error {
	var err error
	DB, err = sql.Open(database, databasePath)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	return nil
}

// CreateTables executes all queries from schema.sql
func CreateTables() error {
	content, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema.sql file: %v", err)
	}

	queries := strings.TrimSpace(string(content))

	_, err = DB.Exec(queries)
	if err != nil {
		return fmt.Errorf("failed to create tables %q: %v", queries, err)
	}

	// insert categories into database if not already exist
	var catCount int
	err = DB.QueryRow(`SELECT COUNT(*) FROM categories`).Scan(&catCount)
	if err != nil {
		return fmt.Errorf("failed to get the count of categories: %v", err)
	}

	if catCount == 0 { // if no categories exist, insert them
		query := `INSERT INTO categories (label) VALUES
            ('Technology'), ('Health'),
            ('Travel'),	('Education'),
            ('Entertainment');`
		_, err = DB.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to insert categories into database: %v", err)
		}
		log.Println("Categories inserted successfully")
	}

	return nil
}

// CreateDemoData generates and inserts fake data into the database
func CreateDemoData() error {
	// create database schema before creating demo data
	if err := CreateTables(); err != nil {
		return err
	}

	// read file that contains all queries to create demo data
	content, err := os.ReadFile(seedPath)
	if err != nil {
		return fmt.Errorf("failed to read seed.sql file: %v", err)
	}

	queries := strings.TrimSpace(string(content))

	_, err = DB.Exec(queries)
	if err != nil {
		log.Printf("failed to insert demo data %q: %v\n", queries, err)
		return err
	}

	log.Println("Demo data created successfully")
	return nil
}
