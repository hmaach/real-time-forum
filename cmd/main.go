package main

import (
	"log"
	"net/http"
	"os"

	"forum/server/config"
	"forum/server/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// check args
	if len(os.Args) != 1 {
		log.Fatalf("Too many arguments")
	}

	// Connect to the database
	db, err := config.Connect()
	if err != nil {
		log.Println("Database connection error:", err)
	}
	defer db.Close()

	err = config.CreateTables(db)
	if err != nil {
		log.Printf("Error creating the database schema: %v\n", err)
	}

	// Start the HTTP server
	server := http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(db),
	}

	log.Println("Server starting on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
