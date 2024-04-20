package main

import (
	"gin-task-api/database"
	"gin-task-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Open connection to in-memory SQLite database
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	// Create a Gin router
	server := gin.Default()

	// Register task router with dependencies
	server = handlers.SetupRoutes(server, db)

	// Start the server on port 8080 (or any desired port)
	server.Run(":8080")
}
