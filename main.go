package main

import (
	apiAuth "chat-app/api/auth"
	apiMessages "chat-app/api/messages"
	"chat-app/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	database.RunMigrations()

	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"localhost"})

	router.POST("/register", apiAuth.Register)
	router.POST("/sign-in", apiAuth.SignIn)

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		println("what is this")
		user := c.MustGet(gin.AuthUserKey).(string)
		println("what is this", user)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	router.POST("/create-message", apiMessages.CreateMessage)
	router.GET("/get-messages", apiMessages.GetMessages)

	router.Run()
}
