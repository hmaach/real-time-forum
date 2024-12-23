package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

// CreateTables executes all queries from schema.sql
func CreateTables(db *sql.DB) error {
	content, err := os.ReadFile(BasePath + "server/database/sql/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema.sql file: %v", err)
	}

	queries := strings.TrimSpace(string(content))

	_, err = db.Exec(queries)
	if err != nil {
		return fmt.Errorf("failed to create tables %q: %v", queries, err)
	}

	// insert categories into database if not already exist
	var catCount int
	err = db.QueryRow(`SELECT COUNT(*) FROM categories`).Scan(&catCount)
	if err != nil {
		return fmt.Errorf("failed to get the count of categories: %v", err)
	}

	if catCount == 0 { // if no categories exist, insert them
		query := `INSERT INTO categories (label) VALUES
			('Technology'), ('Health'),
			('Travel'),	('Education'),
			('Entertainment');`
		_, err = db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to insert categories into database: %v", err)
		}
		log.Println("Categories inserted successfully")
	}

	return nil
}

// CreateFakeData generates and inserts fake data into the database
func CreateDemoData(db *sql.DB) error {
	// create database schema before creating demo data
	if err := CreateTables(db); err != nil {
		return err
	}

	// read file that contains all queries  to create demo data
	content, err := os.ReadFile(BasePath + "server/database/sql/seed.sql")
	if err != nil {
		return fmt.Errorf("failed to read seed.sql file: %v", err)
	}

	queries := strings.TrimSpace(string(content))

	_, err = db.Exec(queries)
	if err != nil {
		log.Printf("failed to isert demo data %q: %v\n", queries, err)
		return err
	}

	log.Println("Demo data created successfully")
	return nil
}
