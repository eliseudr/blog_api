package main

import (
	"fmt"
	"log"

	db "github.com/eliseudr/blog_api/database"
	blog "github.com/eliseudr/blog_api/server"
)

func main() {
	fmt.Println("Blog API is running...")

	// Load and validate all the configuration
	config, err := db.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the database and run migrations
	database, err := db.Initialize(config)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the API server
	server := blog.NewBlogAPIServer(":8080", database)
	log.Fatal(server.Run())
}
