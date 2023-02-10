package messages

import (
	"net/http"

	"chat-app/api/middleware"
	"chat-app/database/models/message"

	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
	messages, err := message.GetMessages()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

type CreateMessageInterface struct {
	Content string `json:"content" binding:"required"`
}

func CreateMessage(c *gin.Context) {
	userId := middleware.GetUserId(c)
	var msgInterface CreateMessageInterface

	if err := c.ShouldBindJSON(&msgInterface); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	println("userId:", userId)

	result, err := message.CreateMessage(&userId, &msgInterface.Content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result.String()})
}
