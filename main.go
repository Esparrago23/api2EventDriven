package main

import (
	"mi-notificacion/src/notifications/infraestructure"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/gin-contrib/cors"
)

func main() {
	r:= gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	infraestructure.NotificationRoutes(r)
	r.Run(":8081")
}
