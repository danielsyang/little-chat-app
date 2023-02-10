package main

import (
	apiAuth "chat-app/api/auth"
	apiMessages "chat-app/api/messages"
	jwtMiddleware "chat-app/api/middleware"
	"chat-app/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.RunMigrations()

	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"localhost"})

	router.POST("/register", apiAuth.Register)
	router.POST("/sign-in", apiAuth.SignIn)

	protected := router.Group("/api")
	protected.Use(jwtMiddleware.JwtAuthMiddleware())

	protected.GET("/get-messages", apiMessages.GetMessages)
	protected.POST("/create-message", apiMessages.CreateMessage)

	router.Run()
}
