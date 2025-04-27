package main

import (
	"log"
	"simple-store-app/config"
	"simple-store-app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize database
	config.InitDB()

	// Routes
	routes.InitRoutes(e)

	// Start server
	log.Fatal(e.Start(":8080"))
}
