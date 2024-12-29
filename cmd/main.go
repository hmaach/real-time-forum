package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"forum/server/api"
	"forum/server/config"
	"forum/server/models"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Parse the HTML index file
	_, err := template.ParseFiles(config.BasePath + "web/index.html")
	if err != nil {
		log.Fatal("Failed to parse index.html:", err)
	}

	// Connect to the database
	err = models.Connect()
	if err != nil {
		log.Panic("Database connection error:", err)
	}
	// err = models.CreateDemoData()
	// if err != nil {
	// 	log.Fatal("error creating demo data:", err)
	// }
}

func main() {
	// check args
	if len(os.Args) != 1 {
		log.Fatal("Too many arguments")
	}

	// Start the HTTP server
	server := http.Server{
		Addr:    ":8080",
		Handler: api.Routes(),
	}

	log.Println("Server starting on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
