package main

import (
	"Backend/config"
	"Backend/controller"
	"Backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnDB()

	Route := gin.Default()

	Route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := Route.Group("/api")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)

		api.POST("/logout", controller.Logout)

		protected := api.Group("/notes")
		protected.Use(middleware.RequireAuth)
		{
			protected.GET("/", controller.GetNotes)
			protected.GET("/:id", controller.GetByID)
			protected.POST("/", controller.CreateNote)
			protected.PUT("/:id", controller.UpdateNote)
			protected.DELETE("/:id", controller.DeleteNote)
		}
	}

	Route.Run(":8080")
}
