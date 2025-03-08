package controllers

import (
	"net/http"

	"mi-notificacion/src/notifications/application"
	"mi-notificacion/src/notifications/domain/entities"

	"github.com/gin-gonic/gin"
)

func NotificationController(c *gin.Context) {
	var notification entities.Notification

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := application.SendNotification(notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to WebSocket server"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification received and sent to WebSocket"})
}
