package main

import (
	"chatting-app/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load environment")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Routes
	routes.RouteInit(e.Group("/api/v1"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Println("Server starting at :8080")
	e.Logger.Fatal(e.Start(":" + port))
}
