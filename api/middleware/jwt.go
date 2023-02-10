package middleware

import (
	"chat-app/services/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		token := strings.Split(bearerToken, " ")[1]
		userId, err := auth.ValidateToken(token)

		if err != nil {
			c.AbortWithStatus(401)
			return ""
		}

		return userId
	}

	c.AbortWithStatus(401)
	return ""
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")

		if len(strings.Split(bearerToken, " ")) != 2 {
			ctx.AbortWithStatus(401)
			return
		}

		token := strings.Split(bearerToken, " ")[1]

		_, err := auth.ValidateToken(token)

		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Next()
	}
}
