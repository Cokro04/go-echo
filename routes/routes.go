package routes

import (
	"simple-store-app/handlers"
	"simple-store-app/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Auth routes
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	// Store routes
	store := e.Group("/stores")
	store.Use(middleware.JWTAuth())
	store.GET("", handlers.GetStores)
	store.POST("", handlers.CreateStore)
	store.PUT("/:id", handlers.UpdateStore)
	store.DELETE("/:id", handlers.DeleteStore)
}