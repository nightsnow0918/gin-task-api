package main

import (
	"gin-task-api/database"
	"gin-task-api/docs"
	"gin-task-api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Open connection to in-memory SQLite database
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	// Create a Gin router
	router := gin.Default()

	// Register task router with dependencies
	router = handlers.SetupRoutes(router, db)

	// Document
	if mode := gin.Mode(); mode == gin.DebugMode {
		docs.SwaggerInfo.BasePath = "/"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Start the router on port 8080 (or any desired port)
	router.Run(":8080")

}
